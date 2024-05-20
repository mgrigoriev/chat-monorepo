package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error) {
	// TODO: Implement real logic

	return &models.User{
		ID:             id,
		Name:           "Updated User",
		Email:          user.Email,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}, nil
}
