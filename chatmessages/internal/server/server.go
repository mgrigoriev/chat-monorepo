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
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases"
	pb "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/api/chatmessages"
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
	cfg                 Config
	Deps
}

func NewServer(ctx context.Context, cfg Config, d Deps) (*Server, error) {
	srv := &Server{
		serverChatMessages:  make(map[uint64]*pb.ChatMessageInfo),
		privateChatMessages: make(map[uint64]*pb.ChatMessageInfo),
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

func (server *Server) Start(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		grpcServer := grpc.NewServer()
		pb.RegisterChatMessagesServiceServer(grpcServer, server)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":"+server.cfg.GrpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		// SaveChatMessage:
		// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/SaveChatMessage '{"info":{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}}'
		// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/SaveChatMessage '{"info":{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}}'
		// ListServerChatMessages:
		// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/ListServerChatMessages '{"server_id":1}'
		// ListPrivateChatMessages:
		// grpc_cli call --json_input --json_output localhost:9090 ChatMessagesService/ListPrivateChatMessages '{"user_id":1, "other_user_id":10}'
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		corsOptions := cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
		}
		corsHandler := cors.New(corsOptions).Handler

		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux := runtime.NewServeMux()
		if err := pb.RegisterChatMessagesServiceHandlerServer(ctx, mux, server); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		httpServer := &http.Server{Handler: corsHandler(mux)}

		lis, err := net.Listen("tcp", ":"+server.cfg.HttpPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		log.Printf("server listening at %v", lis.Addr())
		if err := httpServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		// SaveChatMessage:
		// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}'
		// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}'
		// ListServerChatMessages:
		// curl --location 'localhost:8080/api/v1/chatmessages/server?server_id=1'
		// ListPrivateChatMessages:
		// curl --location 'localhost:8080/api/v1/chatmessages/private?user_id=1&other_user_id=10'
	}()

	wg.Add(1)
	go func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Static("/sw", "./swaggerui")

		e.Logger.Fatal(e.Start(":" + server.cfg.SwaggerPort))
	}()

	wg.Wait()
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
