package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (uc *Usecase) ListServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error) {
	// return nil, models.ErrNotImplemented

	// TODO: Get from repo

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
