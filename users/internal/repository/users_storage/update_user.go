package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) UpdateUser(id models.UserID, user models.User) (*models.User, error) {
	return &models.User{
		ID:             id,
		Name:           "Updated User",
		Email:          user.Email,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}, nil
}
