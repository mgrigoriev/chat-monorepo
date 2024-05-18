package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateInvite(ctx context.Context, chatServerID models.ChatServerID, userID models.UserID) (models.InviteID, error) {
	// return 0, models.ErrNotImplemented
	// TODO: Call repo
	return 20, nil
}
