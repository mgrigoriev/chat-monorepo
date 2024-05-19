package users_storage

import (
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUserByToken(token models.AuthToken) (*models.User, error) {
	if token != "valid-token" {
		return nil, errors.New("invalid token")
	}

	return &models.User{
		ID:             1,
		Name:           "Test 1",
		Email:          "test1@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
