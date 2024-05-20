package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUserByToken(ctx context.Context, token models.AuthToken) (*models.User, error) {
	// TODO: Implement real logic

	if token != "valid-token" {
		return nil, models.ErrDoesNotExist
	}

	return &models.User{
		ID:             1,
		Name:           "Test 1",
		Email:          "test1@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
