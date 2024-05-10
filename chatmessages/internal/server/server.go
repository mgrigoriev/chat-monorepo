package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

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
}

func NewServer() *server {
	return &server{
		serverChatMessages:  make(map[uint64]*pb.ChatMessageInfo),
		privateChatMessages: make(map[uint64]*pb.ChatMessageInfo),
	}
}

func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := NewServer()
	//server, err := NewServer()
	//if err != nil {
	//	log.Fatalf("failed to create server: %v", err)
	//}

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
	}()

	wg.Wait()

	// SaveChatMessage:
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":1, "recipient_id":10, "content":"hello - private message"}'
	// curl --location 'localhost:8080/api/v1/chatmessages' --header 'Content-Type: application/json' --data '{"user_id":1, "user_name":"john", "recipient_type":2, "recipient_id":1, "content":"hello - server message"}'
	// ListServerChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/server?server_id=1'
	// ListPrivateChatMessages:
	// curl --location 'localhost:8080/api/v1/chatmessages/private?user_id=1&other_user_id=10'
}

// SaveChatMessage implements pb.ChatMessagesServiceServer
func (s *server) SaveChatMessage(_ context.Context, req *pb.SaveChatMessageRequest) (*pb.SaveChatMessageResponse, error) {
	info := req.GetInfo()

	log.Printf("SaveChatMessage: received: %s", info.GetContent())

	if err := validateSaveChatMessageRequest(req); err != nil {
		return nil, err
	}

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

func validateSaveChatMessageRequest(req *pb.SaveChatMessageRequest) error {
	info := req.GetInfo()

	var violations []*errdetails.BadRequest_FieldViolation

	if len(info.GetContent()) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "info.content",
			Description: "empty",
		})
	}

	if info.GetUserId() == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "info.user_id",
			Description: "empty",
		})
	}

	if info.GetRecipientId() == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "info.recipient_id",
			Description: "empty",
		})
	}

	if len(info.GetUserName()) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "info.user_name",
			Description: "empty",
		})
	}

	if info.RecipientType != pb.ChatMessageInfo_USER && info.RecipientType != pb.ChatMessageInfo_SERVER {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "info.recipient_type",
			Description: "invalid",
		})
	}

	if len(violations) > 0 {
		st, err := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(&errdetails.BadRequest{
				FieldViolations: violations,
			})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return st.Err()
	}

	return nil
}

// ListServerChatMessages implements pb.ChatMessagesServiceServer
func (s *server) ListServerChatMessages(_ context.Context, req *pb.ListServerChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
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
