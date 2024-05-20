package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) GetUserByID(ctx context.Context, id models.UserID) (*models.User, error) {
	user, err := uc.UsersStorage.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
