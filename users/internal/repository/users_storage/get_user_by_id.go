package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) GetUserByID(id models.UserID) (*models.User, error) {
	return &models.User{
		ID:             id,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}
