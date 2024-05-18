package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateChatServer(ctx context.Context, chatserver models.ChatServer) (models.ChatServerID, error) {
	// return 0, models.ErrNotImplemented
	// TODO: Call repo
	return 1, nil
}
