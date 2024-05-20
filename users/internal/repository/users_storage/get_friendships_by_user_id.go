package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetFriendshipsByUserID(ctx context.Context, userID models.UserID) (*[]models.Friendship, error) {
	// TODO: Implement real logic

	if userID == 2000 {
		return nil, models.ErrDoesNotExist
	}

	return &[]models.Friendship{
		{
			ID:         10,
			FollowerID: 1,
			FollowedID: 2,
			Status:     "approved",
		},
	}, nil
}
