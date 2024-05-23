package chatservers_storage

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatservers/pkg/errors"
)

func (r *ChatServersStorage) CreateChatServer(ctx context.Context, chatServer models.ChatServer) (models.ChatServerID, error) {
	const api = "chatservers_storage.CreateChatServer"

	row, err := newChatServerRowFromModel(&chatServer)
	if err != nil {
		return 0, pkgerrors.Wrap(api, err)
	}

	query := squirrel.Insert(chatserversTable).
		SetMap(row.ValuesMap()).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id")

	var id models.ChatServerID
	if err := r.driver.GetQueryEngine(ctx).Getx(ctx, &id, query); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return 0, pkgerrors.Wrap(api, models.ErrAlreadyExists)
		}
		return 0, pkgerrors.Wrap(api, err)
	}

	return id, nil
}
