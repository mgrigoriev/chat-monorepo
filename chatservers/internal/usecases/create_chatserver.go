package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateChatServer(ctx context.Context, chatServer models.ChatServer) (models.ChatServerID, error) {
	chatServerID, err := uc.ChatServersStorage.CreateChatServer(ctx, chatServer)
	if err != nil {
		return 0, err
	}

	return chatServerID, nil
}
