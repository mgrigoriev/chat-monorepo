package chatservers_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatservers/pkg/errors"
)

func (r *ChatServersStorage) GetChatServersByUserID(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error) {
	const api = "chatservers_storage.GetChatServersByUserID"

	query := squirrel.Select("id", "user_id", "name").
		From(chatserversTable).
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar)

	rows := make([]chatServerRow, 0)
	if err := r.driver.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	chatServers := make([]models.ChatServer, 0, len(rows))
	for _, row := range rows {
		chatServers = append(chatServers, models.ChatServer{
			ID:     models.ChatServerID(row.ID),
			UserID: models.UserID(row.UserID),
			Name:   row.Name,
		})
	}

	return &chatServers, nil
}
