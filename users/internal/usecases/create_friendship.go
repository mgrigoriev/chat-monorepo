package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) CreateFriendship(ctx context.Context, friendship models.Friendship) (models.FriendshipID, error) {
	friendship.Status = "pending"
	friendshipID, err := uc.UsersStorage.CreateFriendship(ctx, friendship)
	if err != nil {
		return 0, err
	}

	return friendshipID, nil
}
