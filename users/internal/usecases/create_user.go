package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) CreateUser(ctx context.Context, user models.User) (models.UserID, error) {
	userID, err := uc.UsersStorage.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
