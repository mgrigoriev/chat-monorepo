package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) SearchUsers(ctx context.Context, term string) (*[]models.User, error) {
	//return nil, models.ErrNotImplemented
	// TODO: Call repo
	return &[]models.User{
		{
			ID:             1,
			Name:           "Test",
			Email:          "test@test.com",
			AvatarPhotoURL: "https://test.com/1.jpg",
		},
	}, nil
}
