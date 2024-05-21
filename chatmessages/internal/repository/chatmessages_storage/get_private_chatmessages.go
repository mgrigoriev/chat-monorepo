package chatmessages_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (r *ChatMessagesStorage) GetPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error) {
	// TODO: Implement real logic
	_ = userID
	_ = otherUserID

	return &[]models.ChatMessage{
		{
			ID:            1,
			UserID:        2,
			UserName:      "Author Name",
			RecipientType: 1,
			RecipientID:   3,
			Content:       "test private message",
		},
	}, nil
}
