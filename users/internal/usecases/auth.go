package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) Auth(ctx context.Context, token models.AuthToken) (*models.User, error) {
	//return nil, models.ErrNotImplemented

	// TODO: Call repo

	if token != "valid-token" {
		return nil, errors.New("invalid token")
	}

	user, err := uc.Repo.GetUserByToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
