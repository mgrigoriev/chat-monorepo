package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) Auth(ctx context.Context, token models.AuthToken) (*models.User, error) {
	user, err := uc.UsersStorage.GetUserByToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
