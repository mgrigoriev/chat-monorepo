package server

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/api/chatmessages"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var idSerial uint64

// server is used to implement pb.ChatMessagesServiceServer.
type server struct {
	// UnimplementedChatMessagesServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedChatMessagesServiceServer
	mx                  sync.RWMutex
	serverChatMessages  map[uint64]*pb.ChatMessageInfo
	privateChatMessages map[uint64]*pb.ChatMessageInfo
	validator           *protovalidate.Validator
}

func NewServer() (*server, error) {
	srv := &server{
		serverChatMessages:  make(map[uint64]*pb.ChatMessageInfo),
		privateChatMessages: make(map[uint64]*pb.ChatMessageInfo),
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

func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server, err := NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		grpcServer := grpc.NewServer()
		pb.RegisterChatMessagesServiceServer(grpcServer, server)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":9090")
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
		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux := runtime.NewServeMux()
		if err := pb.RegisterChatMessagesServiceHandlerServer(ctx, mux, server); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		httpServer := &http.Server{Handler: mux}

		lis, err := net.Listen("tcp", ":8080")
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

// SaveChatMessage implements pb.ChatMessagesServiceServer
func (s *server) SaveChatMessage(_ context.Context, req *pb.SaveChatMessageRequest) (*pb.SaveChatMessageResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	info := req.GetInfo()

	log.Printf("SaveChatMessage: received: %s", info.GetContent())

	id := atomic.AddUint64(&idSerial, 1)

	s.mx.Lock()
	if info.RecipientType == pb.ChatMessageInfo_SERVER {
		s.serverChatMessages[id] = info
	} else {
		s.privateChatMessages[id] = info
	}
	s.mx.Unlock()

	return &pb.SaveChatMessageResponse{
		Id: id,
	}, nil
}

// ListServerChatMessages implements pb.ChatMessagesServiceServer
func (s *server) ListServerChatMessages(_ context.Context, req *pb.ListServerChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	serverID := req.GetServerId()

	log.Println("ListServerChatMessages: received")

	s.mx.RLock()
	defer s.mx.RUnlock()

	chatMessages := make([]*pb.ChatMessage, 0, len(s.serverChatMessages))
	for id, msg := range s.serverChatMessages {
		if msg.RecipientId == serverID {
			chatMessages = append(chatMessages, &pb.ChatMessage{
				Id:   id,
				Info: msg,
			})
		}
	}

	return &pb.ListChatMessagesResponse{
		ChatMessages: chatMessages,
	}, nil
}

func (s *server) ListPrivateChatMessages(_ context.Context, req *pb.ListPrivateChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	userID := req.UserId
	otherUserID := req.OtherUserId

	log.Println("ListPrivateChatMessages: received")

	s.mx.RLock()
	defer s.mx.RUnlock()

	chatMessages := make([]*pb.ChatMessage, 0, len(s.privateChatMessages))
	for id, msg := range s.privateChatMessages {
		if (msg.UserId == userID && msg.RecipientId == otherUserID) || (msg.UserId == otherUserID && msg.RecipientId == userID) {
			chatMessages = append(chatMessages, &pb.ChatMessage{
				Id:   id,
				Info: msg,
			})
		}
	}

	return &pb.ListChatMessagesResponse{
		ChatMessages: chatMessages,
	}, nil
}
