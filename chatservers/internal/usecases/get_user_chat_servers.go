package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) GetUserChatServers(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error) {
	chatServers, err := uc.ChatServersStorage.GetChatServersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return chatServers, nil
}
