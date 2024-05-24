package users_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) DeleteFriendship(ctx context.Context, friendshipID models.FriendshipID) error {
	const api = "users_storage.DeleteFriendship"

	query := squirrel.Delete(friendshipsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": friendshipID})

	if _, err := r.driver.GetQueryEngine(ctx).Execx(ctx, query); err != nil {
		return pkgerrors.Wrap(api, err)
	}

	return nil
}
