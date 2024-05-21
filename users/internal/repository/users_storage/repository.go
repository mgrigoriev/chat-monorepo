package users_storage

import uc "github.com/mgrigoriev/chat-monorepo/users/internal/usecases"

// Check that we implement contract for usecase
var (
	_ uc.UsersStorage = (*UsersStorage)(nil)
)

type UsersStorage struct{}

func New() *UsersStorage {
	return &UsersStorage{}
}
