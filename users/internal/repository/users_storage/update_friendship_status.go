package users_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) UpdateFriendshipStatus(ctx context.Context, friendshipID models.FriendshipID, status string) error {
	const api = "users_storage.UpdateFriendshipStatus"

	query := squirrel.Update(friendshipsTable).
		Set("status", status).
		Where(squirrel.Eq{"id": friendshipID}).
		PlaceholderFormat(squirrel.Dollar)

	result, err := r.driver.GetQueryEngine(ctx).Execx(ctx, query)
	if err != nil {
		return pkgerrors.Wrap(api, err)
	}

	if result.RowsAffected() == 0 {
		return pkgerrors.Wrap(api, models.ErrDoesNotExist)
	}

	return nil
}
