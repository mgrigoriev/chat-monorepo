package users_storage

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) CreateUser(ctx context.Context, user models.User) (models.UserID, error) {
	// TODO: Implement real logic

	if user.Name == "Existing User" {
		return models.UserID(0), models.ErrAlreadyExists
	}

	return models.UserID(1), nil
}
