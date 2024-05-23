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

func (r *ChatServersStorage) CreateParticipant(ctx context.Context, chatServer models.Participant) (models.ParticipantID, error) {
	const api = "chatservers_storage.CreateParticipant"

	row, err := newParticipantRowFromModel(&chatServer)
	if err != nil {
		return 0, pkgerrors.Wrap(api, err)
	}

	query := squirrel.Insert(participantsTable).
		SetMap(row.ValuesMap()).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id")

	var id models.ParticipantID
	if err := r.driver.GetQueryEngine(ctx).Getx(ctx, &id, query); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return 0, pkgerrors.Wrap(api, models.ErrAlreadyExists)
		}
		return 0, pkgerrors.Wrap(api, err)
	}

	return id, nil
}
