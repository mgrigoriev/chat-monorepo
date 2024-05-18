package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error) {
	//return nil, models.ErrNotImplemented
	// TODO: Call repo
	return &models.User{
		ID:             30,
		Name:           "Updated User",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
