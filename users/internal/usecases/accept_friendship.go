package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) AcceptFriendship(ctx context.Context, friendshipID models.FriendshipID) error {
	err := uc.UsersStorage.UpdateFriendshipStatus(friendshipID, "accepted")
	if err != nil {
		return err
	}

	return nil
}
