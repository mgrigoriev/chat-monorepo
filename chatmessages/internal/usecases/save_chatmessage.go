package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/models"
)

func (uc *Usecase) SaveChatMessage(ctx context.Context, chatMessage models.ChatMessage) (models.ChatMessageID, error) {
	chatMessageID, err := uc.ChatMessagesStorage.CreateChatMessage(ctx, chatMessage)
	if err != nil {
		return 0, err
	}

	return chatMessageID, nil
}
