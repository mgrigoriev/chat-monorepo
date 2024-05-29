package chatmessages_storage

import (
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/models"
)

type chatMessageRow struct {
	ID            int64  `db:"id"`
	UserID        int64  `db:"user_id"`
	UserName      string `db:"user_name"`
	RecipientType int8   `db:"recipient_type"`
	RecipientID   int64  `db:"recipient_id"`
	Content       string `db:"content"`
}

func (r *chatMessageRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":             r.ID,
		"user_id":        r.UserID,
		"user_name":      r.UserName,
		"recipient_type": r.RecipientType,
		"recipient_id":   r.RecipientID,
		"content":        r.Content,
	}
}

func (r *chatMessageRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func newChatMessageRowFromModel(chatMessage *models.ChatMessage) (*chatMessageRow, error) {
	return &chatMessageRow{
		ID:            int64(chatMessage.ID),
		UserID:        int64(chatMessage.UserID),
		UserName:      chatMessage.UserName,
		RecipientType: int8(chatMessage.RecipientType),
		RecipientID:   int64(chatMessage.RecipientID),
		Content:       chatMessage.Content,
	}, nil
}
