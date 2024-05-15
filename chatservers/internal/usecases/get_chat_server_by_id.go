package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error) {
	//return nil, models.ErrNotImplemented
	return &models.ChatServer{
		ID:     1,
		UserID: 2,
		Name:   "Chatserver example",
	}, nil
}
