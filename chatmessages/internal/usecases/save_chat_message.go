package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (uc *Usecase) SaveChatMessage(ctx context.Context, chatmessage models.ChatMessage) (models.ChatMessageID, error) {
	// return 0, models.ErrNotImplemented

	// TODO: Save to repo

	return 1, nil
}
