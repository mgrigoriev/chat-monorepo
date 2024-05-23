package chatservers_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatservers/pkg/errors"
)

func (r *ChatServersStorage) GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error) {
	const api = "chatservers_storage.GetChatServerByID"

	query := squirrel.Select("id", "user_id", "name").
		From(chatserversTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	var row chatServerRow
	err := r.driver.GetQueryEngine(ctx).Getx(ctx, &row, query)
	if err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	chatServer := models.ChatServer{
		ID:     models.ChatServerID(row.ID),
		UserID: models.UserID(row.UserID),
		Name:   row.Name,
	}

	return &chatServer, nil
}
