package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error) {
	updatedUser, err := uc.UsersStorage.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
