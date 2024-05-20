package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) UpdateFriendshipStatus(ctx context.Context, friendshipID models.FriendshipID, status string) error {
	// TODO: Implement real logic

	if friendshipID == 2000 {
		return models.ErrDoesNotExist
	}

	return nil
}
