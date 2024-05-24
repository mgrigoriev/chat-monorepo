package users_storage

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) CreateUser(ctx context.Context, user models.User) (models.UserID, error) {
	const api = "users_storage.CreateUser"

	row, err := newUserRowFromModel(&user)
	if err != nil {
		return 0, pkgerrors.Wrap(api, err)
	}

	query := squirrel.Insert(usersTable).
		SetMap(row.ValuesMap()).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id")

	var id models.UserID
	if err := r.driver.GetQueryEngine(ctx).Getx(ctx, &id, query); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return 0, pkgerrors.Wrap(api, models.ErrAlreadyExists)
		}
		return 0, pkgerrors.Wrap(api, err)
	}

	return id, nil
}
