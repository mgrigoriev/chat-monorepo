package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) DeleteFriendship(ctx context.Context, friendshipID models.FriendshipID) error {
	// TODO: Implement real logic

	if friendshipID == 2000 {
		return models.ErrDoesNotExist
	}

	return nil
}
