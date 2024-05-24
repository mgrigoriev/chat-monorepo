package users_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) GetFriendshipsByUserID(ctx context.Context, userID models.UserID) (*[]models.Friendship, error) {
	const api = "users_storage.GetFriendshipsByUserID"

	query := squirrel.Select("id", "follower_id", "followed_id").
		From(friendshipsTable).
		Where(squirrel.Or{
			squirrel.Eq{"follower_id": userID},
			squirrel.Eq{"followed_id": userID},
		}).
		PlaceholderFormat(squirrel.Dollar)

	rows := make([]friendshipRow, 0)
	if err := r.driver.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	friendships := make([]models.Friendship, 0, len(rows))
	for _, row := range rows {
		friendships = append(
			friendships,
			models.Friendship{
				ID:         models.FriendshipID(row.ID),
				FollowerID: models.UserID(row.FollowerID),
				FollowedID: models.UserID(row.FollowedID),
				Status:     row.Status,
			},
		)
	}

	return &friendships, nil
}
