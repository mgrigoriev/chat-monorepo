package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) SearchUsers(ctx context.Context, term string) (*[]models.User, error) {
	users, err := uc.UsersStorage.GetUsersByTerm(ctx, term)
	if err != nil {
		return nil, err
	}

	return users, nil
}
