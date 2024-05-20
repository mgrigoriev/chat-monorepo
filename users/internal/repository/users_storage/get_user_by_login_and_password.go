package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUserByLoginAndPassword(ctx context.Context, login string, password string) (*models.User, error) {
	// TODO: Implement real logic

	if login == "cool" && password == "hacker" {
		return nil, models.ErrDoesNotExist
	}

	return &models.User{
		ID:             3,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
