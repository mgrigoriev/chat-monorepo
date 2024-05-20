package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) CreateFriendship(ctx context.Context, followerID models.UserID, followedID models.UserID) (models.FriendshipID, error) {
	// TODO: Implement real logic

	if followerID == 1000 && followedID == 1001 {
		return models.FriendshipID(0), models.ErrAlreadyExists
	}

	return models.FriendshipID(1), nil
}
