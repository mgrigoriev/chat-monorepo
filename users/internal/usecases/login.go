package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) Login(ctx context.Context, email string, password string) (models.AuthToken, error) {
	_, err := uc.UsersStorage.GetUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return "", err
	}

	// TODO: Update logic
	return "valid-token", nil
}
