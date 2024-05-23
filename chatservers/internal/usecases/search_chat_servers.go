package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) SearchChatServers(ctx context.Context, term string) (*[]models.ChatServer, error) {
	chatServers, err := uc.ChatServersStorage.GetChatServersByTerm(ctx, term)
	if err != nil {
		return nil, err
	}

	return chatServers, nil
}
