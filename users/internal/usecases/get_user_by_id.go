package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) GetUserByID(ctx context.Context, id models.UserID) (*models.User, error) {
	//return nil, models.ErrNotImplemented
	// TODO: Call repo
	return &models.User{
		ID:             1,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
