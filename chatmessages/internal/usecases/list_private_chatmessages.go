package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

func (uc *Usecase) ListPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error) {
	chatMessages, err := uc.ChatMessagesStorage.GetPrivateChatMessages(ctx, userID, otherUserID)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}
