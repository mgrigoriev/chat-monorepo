package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/models"
)

func (uc *Usecase) ListServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error) {
	chatMessages, err := uc.ChatMessagesStorage.GetServerChatMessages(ctx, serverID)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}
