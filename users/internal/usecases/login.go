package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (uc *Usecase) Login(ctx context.Context, login string, password string) (models.AuthToken, error) {
	//return "", models.ErrNotImplemented
	// TODO: Call repo
	return "valid-token", nil
}
