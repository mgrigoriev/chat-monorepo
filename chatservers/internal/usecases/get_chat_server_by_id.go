package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error) {
	chatServer, err := uc.ChatServersStorage.GetChatServerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return chatServer, nil
}
