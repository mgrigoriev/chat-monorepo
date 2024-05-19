package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) GetUsersByNameSubstring(nameSubstring string) (*[]models.User, error) {
	_ = nameSubstring

	return &[]models.User{
		{
			ID:             1,
			Name:           "Test",
			Email:          "test@test.com",
			AvatarPhotoURL: "https://test.com/1.jpg",
		},
	}, nil
}
