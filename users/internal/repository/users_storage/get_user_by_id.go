package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUserByID(ctx context.Context, id models.UserID) (*models.User, error) {
	// TODO: Implement real logic

	if id == 2000 {
		return nil, models.ErrDoesNotExist
	}

	return &models.User{
		ID:             id,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
