package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) CreateUser(user models.User) (models.UserID, error) {
	_ = user

	return models.UserID(1), nil
}
