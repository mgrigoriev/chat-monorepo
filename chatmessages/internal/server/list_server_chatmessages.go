package server

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	pb "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/api/chatmessages"
	"log"
)

// ListServerChatMessages implements pb.ChatMessagesServiceServer
func (s *Server) ListServerChatMessages(ctx context.Context, req *pb.ListServerChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	serverID := req.GetServerId()

	messages, err := s.Usecase.ListServerChatMessages(ctx, models.ChatServerID(serverID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform gRPC request: %w", err)
	}

	log.Println("ListServerChatMessages: received")

	chatMessages := make([]*pb.ChatMessage, 0, len(s.privateChatMessages))
	for _, msg := range *messages {
		chatMessages = append(chatMessages, &pb.ChatMessage{
			Id: uint64(msg.ID),
			Info: &pb.ChatMessageInfo{
				UserId:        uint64(msg.UserID),
				UserName:      msg.UserName,
				RecipientType: pb.ChatMessageInfo_RECIPIENT_TYPE(msg.RecipientType),
				RecipientId:   uint64(msg.RecipientID),
				Content:       msg.Content,
			},
		})
	}

	return &pb.ListChatMessagesResponse{
		ChatMessages: chatMessages,
	}, nil
}
