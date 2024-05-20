package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUsersByNameSubstring(ctx context.Context, nameSubstring string) (*[]models.User, error) {
	// TODO: Implement real logic

	return &[]models.User{
		{
			ID:             1,
			Name:           "Test",
			Email:          "test@test.com",
			AvatarPhotoURL: "https://test.com/1.jpg",
		},
	}, nil
}
