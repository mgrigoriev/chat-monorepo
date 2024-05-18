package server

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	pb "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/api/chatmessages"
	"log"
)

func (s *Server) ListPrivateChatMessages(ctx context.Context, req *pb.ListPrivateChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	userID := req.UserId
	otherUserID := req.OtherUserId

	messages, err := s.Usecase.ListPrivateChatMessages(ctx, models.UserID(userID), models.UserID(otherUserID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform gRPC request: %w", err)
	}

	log.Println("ListPrivateChatMessages: received")

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
