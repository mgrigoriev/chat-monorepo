package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) CreateFriendship(ctx context.Context, followerID models.UserID, followedID models.UserID) (models.FriendshipID, error) {
	friendshipID, err := uc.UsersStorage.CreateFriendship(followerID, followedID)
	if err != nil {
		return 0, err
	}

	return friendshipID, nil
}
