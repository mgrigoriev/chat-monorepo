package chatmessages_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (r *ChatMessagesStorage) GetServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error) {
	// TODO: Implement real logic
	_ = serverID

	return &[]models.ChatMessage{
		{
			ID:            1,
			UserID:        2,
			UserName:      "Author Name",
			RecipientType: 1,
			RecipientID:   3,
			Content:       "test server message",
		},
	}, nil
}
