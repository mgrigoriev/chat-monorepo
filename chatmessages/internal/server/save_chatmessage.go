package server

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	pb "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/api/chatmessages"
	"log"
)

// SaveChatMessage implements pb.ChatMessagesServiceServer
func (s *Server) SaveChatMessage(ctx context.Context, req *pb.SaveChatMessageRequest) (*pb.SaveChatMessageResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	info := req.GetInfo()

	log.Printf("SaveChatMessage: received: %s", info.GetContent())

	chatMessage := models.ChatMessage{
		UserID:        models.UserID(info.UserId),
		UserName:      info.UserName,
		RecipientType: models.RecipientType(info.RecipientType),
		RecipientID:   models.UserID(info.RecipientId),
		Content:       info.Content,
	}

	id, err := s.Usecase.SaveChatMessage(ctx, chatMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to perform gRPC request: %w", err)
	}

	return &pb.SaveChatMessageResponse{
		Id: uint64(id),
	}, nil
}
