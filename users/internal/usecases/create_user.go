package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) CreateUser(ctx context.Context, user models.User) (models.UserID, error) {
	//return 0, models.ErrNotImplemented

	return 20, nil
}
