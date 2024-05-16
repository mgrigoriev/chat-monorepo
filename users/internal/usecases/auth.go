package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) Auth(ctx context.Context, token models.AuthToken) (*models.User, error) {
	//return nil, models.ErrNotImplemented

	if token != "valid-token" {
		return nil, errors.New("invalid token")
	}

	return &models.User{
		ID:             1,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "http://test.com/1.jpg",
	}, nil
}
