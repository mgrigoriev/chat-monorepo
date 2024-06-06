package server

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases"
	pb "github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/api/chatmessages"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"sync"
	"time"
)

const (
	certFilePath = "./cert/server-cert.crt" // public (клиент шифрует + убеждается что нам можно доверять)
	keyFilePath  = "./cert/server-key.key"  // private key (расшифровываем)
)

var idSerial uint64

type Config struct {
	GrpcPort               string
	HttpPort               string
	InternalServerPort     string
	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Deps struct {
	Usecase usecases.UsecaseInterface
}

// Server is used to implement pb.ChatMessagesServiceServer.
type Server struct {
	// UnimplementedChatMessagesServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedChatMessagesServiceServer
	mx                  sync.RWMutex
	serverChatMessages  map[uint64]*pb.ChatMessageInfo
	privateChatMessages map[uint64]*pb.ChatMessageInfo
	validator           *protovalidate.Validator
	grpcServer          *grpc.Server
	mux                 *runtime.ServeMux
	httpServer          *http.Server
	internalServer      *echo.Echo
	cfg                 Config
	Deps
}

func NewServer(cfg Config, d Deps) (*Server, error) {
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}
	corsHandler := cors.New(corsOptions).Handler
	mux := runtime.NewServeMux()
	httpServer := http.Server{Handler: corsHandler(mux)}

	grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
	grpcServerOptions = append(grpcServerOptions, grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...))

	if os.Getenv("TLS") == "true" {
		tlsConfig, err := createServerTLSConfig(certFilePath, keyFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		grpcServerOptions = append(grpcServerOptions, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	grpcServer := grpc.NewServer(grpcServerOptions...)

	srv := &Server{
		serverChatMessages:  make(map[uint64]*pb.ChatMessageInfo),
		privateChatMessages: make(map[uint64]*pb.ChatMessageInfo),
		grpcServer:          grpcServer,
		mux:                 mux,
		httpServer:          &httpServer,
		internalServer:      echo.New(),
		Deps:                d,
		cfg:                 cfg,
	}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.SaveChatMessageRequest{},
			&pb.ListPrivateChatMessagesRequest{},
			&pb.ListServerChatMessagesRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator
	return srv, nil
}

func (s *Server) Start(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startGRPCServer(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startHTTPGatewayServer(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startInternalServer(ctx)
	}()

	go func() {
		// Wait until we receive a shutdown signal
		<-ctx.Done()
		s.gracefulShutdown(ctx)
	}()

	wg.Wait()
}

func (s *Server) startGRPCServer(ctx context.Context) {
	pb.RegisterChatMessagesServiceServer(s.grpcServer, s)

	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", ":"+s.cfg.GrpcPort)
	if err != nil {
		logger.Fatalf(ctx, "grpc server: failed to listen: %v", err)
	}

	logger.Infof(ctx, "grpc server: server listening at %v", lis.Addr())
	if err := s.grpcServer.Serve(lis); err != nil {
		logger.Infof(ctx, "grpc server: failed to serve: %v", err)
	}

	// SaveChatMessage:
	// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/SaveChatMessage '{"info":{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}}'
	// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/SaveChatMessage '{"info":{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}}'
	// ListServerChatMessages:
	// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/ListServerChatMessages '{"server_id":1}'
	// ListPrivateChatMessages:
	// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/ListPrivateChatMessages '{"user_id":1, "other_user_id":10}'
}

func (s *Server) startHTTPGatewayServer(ctx context.Context) {
	if err := pb.RegisterChatMessagesServiceHandlerServer(ctx, s.mux, s); err != nil {
		logger.Fatalf(ctx, "http gateway: failed to serve: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+s.cfg.HttpPort)
	if err != nil {
		logger.Fatalf(ctx, "http gateway: failed to listen: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	logger.Infof(ctx, "http gateway: server listening at %v", lis.Addr())
	if err := s.httpServer.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Infof(ctx, "http gateway: failed to serve: %v", err)
	}

	// SaveChatMessage:
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}'
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}'
	// ListServerChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/server?server_id=1'
	// ListPrivateChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/private?user_id=1&other_user_id=10'
}

func (s *Server) startInternalServer(ctx context.Context) {
	s.internalServer.Use(middleware.Logger())
	s.internalServer.Static("/sw", "./swaggerui")

	// prometheus metrics
	s.internalServer.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// pprof
	s.internalServer.GET("/debug/pprof/", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	s.internalServer.GET("/debug/pprof/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	s.internalServer.GET("/debug/pprof/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	s.internalServer.GET("/debug/pprof/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	s.internalServer.GET("/debug/pprof/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))

	if err := s.internalServer.Start(":" + s.cfg.InternalServerPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Infof(ctx, "internal server: failed to serve: %v", err)
	}
}

func (s *Server) gracefulShutdown(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof(ctx, "internal server: shutting down server gracefully")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.internalServer.Shutdown(ctx); err != nil {
			s.internalServer.Close()
			logger.Infof(ctx, "internal server: shut down error: %v", err)
		} else {
			log.Print("internal server: shut down")
		}
	}()

	wg.Wait() // Ensure Swagger server is shutdown first

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Print("http gateway: shutting down server gracefully")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.httpServer.Shutdown(ctx); err != nil {
			s.httpServer.Close()
			logger.Infof(ctx, "http gateway: shut down error: %v", err)
		} else {
			log.Print("http gateway: shut down")
		}
	}()

	wg.Wait() // Ensure HTTP Gateway server is shutdown second

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Print("grpc server: shutting down server gracefully")
		s.grpcServer.GracefulStop()
		log.Print("grpc server: shut down")
	}()

	wg.Wait() // Ensure gRPC server is shutdown last
}

func protovalidateViolationsToGoogleViolations(vs []*validate.Violation) []*errdetails.BadRequest_FieldViolation {
	res := make([]*errdetails.BadRequest_FieldViolation, len(vs))
	for i, v := range vs {
		res[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.FieldPath,
			Description: v.Message,
		}
	}
	return res
}

func convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr *protovalidate.ValidationError) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: protovalidateViolationsToGoogleViolations(valErr.Violations),
	}
}

func rpcValidationError(err error) error {
	if err == nil {
		return nil
	}

	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		st, err := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr))
		if err == nil {
			return st.Err()
		}
	}

	return status.Error(codes.Internal, err.Error())
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}

func createServerTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	// Load server's certificate and private key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load x509: %v", err)
	}

	// Create tls config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	}

	return tlsConfig, nil
}
