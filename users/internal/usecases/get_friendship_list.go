package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) GetFriendshipList(ctx context.Context, userID models.UserID) (*[]models.Friendship, error) {
	//return nil, models.ErrNotImplemented
	return &[]models.Friendship{
		{
			ID:         10,
			FollowerID: 1,
			FollowedID: 2,
			Status:     "approved",
		},
	}, nil
}
