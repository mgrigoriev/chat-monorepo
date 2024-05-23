package chatmessages_storage

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/errors"
)

func (r *ChatMessagesStorage) CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) (models.ChatMessageID, error) {
	const api = "chatmessages_storage.CreateChatMessage"

	row, err := newChatMessageRowFromModel(&chatMessage)
	if err != nil {
		return 0, pkgerrors.Wrap(api, err)
	}

	query := squirrel.Insert(chatmessagesTable).
		SetMap(row.ValuesMap()).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id")

	var id models.ChatMessageID
	if err := r.driver.GetQueryEngine(ctx).Getx(ctx, &id, query); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return 0, pkgerrors.Wrap(api, models.ErrAlreadyExists)
		}
		return 0, pkgerrors.Wrap(api, err)
	}

	return id, nil
}
