package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) DeclineFriendship(ctx context.Context, friendshipID models.FriendshipID) error {
	err := uc.UsersStorage.UpdateFriendshipStatus(ctx, friendshipID, "declined")
	if err != nil {
		return err
	}

	return nil
}
