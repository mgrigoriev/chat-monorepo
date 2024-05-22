package chatmessages_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/errors"
)

func (r *ChatMessagesStorage) GetServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error) {
	const api = "chatmessages_storage.GetPrivateChatMessages"

	query := squirrel.Select("*").
		From("chatmessages").
		Where(squirrel.Eq{"recipient_type": 2, "recipient_id": serverID}).
		PlaceholderFormat(squirrel.Dollar)

	rows := make([]chatMessageRow, 0)
	err := r.driver.GetQueryEngine(ctx).Selectx(ctx, &rows, query)
	if err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	chatMessages := make([]models.ChatMessage, 0)
	for _, row := range rows {
		chatMessage := models.ChatMessage{
			ID:            models.ChatMessageID(row.ID),
			UserID:        models.UserID(row.UserID),
			UserName:      row.UserName,
			RecipientType: models.RecipientType(row.RecipientType),
			RecipientID:   models.UserID(row.RecipientID),
			Content:       row.Content,
		}

		chatMessages = append(chatMessages, chatMessage)

	}

	return &chatMessages, nil
}
