package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) GetUserByLoginAndPassword(login string, password string) (*models.User, error) {
	_ = login
	_ = password

	return &models.User{
		ID:             3,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
