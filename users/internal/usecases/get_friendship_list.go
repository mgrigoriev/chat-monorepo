package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) GetFriendshipList(ctx context.Context, userID models.UserID) (*[]models.Friendship, error) {
	friendships, err := uc.UsersStorage.GetFriendshipsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return friendships, nil
}
