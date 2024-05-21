package chatmessages_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (r *ChatMessagesStorage) CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) (models.ChatMessageID, error) {
	// TODO: Implement real logic
	_ = chatMessage

	return models.ChatMessageID(1), nil
}
