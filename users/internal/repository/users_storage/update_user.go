package users_storage

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
	"github.com/pkg/errors"
)

func (r *UsersStorage) UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error) {
	const api = "users_storage.UpdateUser"

	row, err := newUserRowFromModel(&user)
	if err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	query := squirrel.Update(usersTable).
		SetMap(row.ValuesMap()).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING *")

	var updatedUser models.User
	if err := r.driver.GetQueryEngine(ctx).Getx(ctx, &updatedUser, query); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return nil, pkgerrors.Wrap(api, models.ErrAlreadyExists)
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil, pkgerrors.Wrap(api, models.ErrDoesNotExist)
		}
		return nil, pkgerrors.Wrap(api, err)
	}

	return &updatedUser, nil
}
