package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (uc *Usecase) ListPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error) {
	// return nil, models.ErrNotImplemented

	// TODO: Get from repo

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
