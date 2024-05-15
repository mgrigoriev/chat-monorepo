package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) GetUserChatServers(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error) {
	// return nil, models.ErrNotImplemented
	return &[]models.ChatServer{
		{
			ID:     1,
			UserID: 2,
			Name:   "Chatserver example 1",
		},
		{
			ID:     2,
			UserID: 3,
			Name:   "Chatserver example 2",
		},
	}, nil
}
