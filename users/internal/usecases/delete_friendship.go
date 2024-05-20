package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) DeleteFriendship(ctx context.Context, friendshipID models.FriendshipID) error {
	err := uc.UsersStorage.DeleteFriendship(ctx, friendshipID)
	if err != nil {
		return err
	}

	return nil
}
