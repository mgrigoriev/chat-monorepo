package server

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"context"
	"errors"
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases"
	pb "github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/api/chatmessages"
	"github.com/rs/cors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var idSerial uint64

type Config struct {
	GrpcPort    string
	HttpPort    string
	SwaggerPort string
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
	swaggerServer       *echo.Echo
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

	srv := &Server{
		serverChatMessages:  make(map[uint64]*pb.ChatMessageInfo),
		privateChatMessages: make(map[uint64]*pb.ChatMessageInfo),
		grpcServer:          grpc.NewServer(),
		mux:                 mux,
		httpServer:          &http.Server{Handler: corsHandler(mux)},
		swaggerServer:       echo.New(),
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
		s.startSwaggerServer(ctx)
	}()

	go func() {
		// Wait until we receive a shutdown signal
		<-ctx.Done()
		s.gracefulShutdown()
	}()

	wg.Wait()
}

func (s *Server) startGRPCServer(ctx context.Context) {
	pb.RegisterChatMessagesServiceServer(s.grpcServer, s)

	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", ":"+s.cfg.GrpcPort)
	if err != nil {
		log.Fatalf("grpc server: failed to listen: %v", err)
	}

	log.Printf("grpc server: server listening at %v", lis.Addr())
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Printf("grpc server: failed to serve: %v", err)
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
		log.Fatalf("http gateway: failed to serve: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+s.cfg.HttpPort)
	if err != nil {
		log.Fatalf("http gateway: failed to listen: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("http gateway: server listening at %v", lis.Addr())
	if err := s.httpServer.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("http gateway: failed to serve: %v", err)
	}

	// SaveChatMessage:
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}'
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}'
	// ListServerChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/server?server_id=1'
	// ListPrivateChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/private?user_id=1&other_user_id=10'
}

func (s *Server) startSwaggerServer(ctx context.Context) {
	s.swaggerServer.Use(middleware.Logger())
	s.swaggerServer.Static("/sw", "./swaggerui")

	if err := s.swaggerServer.Start(":" + s.cfg.SwaggerPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("swagger server: failed to serve: %v", err)
	}
}

func (s *Server) gracefulShutdown() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("swagger server: shutting down server gracefully")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.swaggerServer.Shutdown(ctx); err != nil {
			s.swaggerServer.Close()
			log.Printf("swagger server: shut down error: %v", err)
		} else {
			log.Print("swagger server: shut down")
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
			log.Printf("http gateway: shut down error: %v", err)
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
