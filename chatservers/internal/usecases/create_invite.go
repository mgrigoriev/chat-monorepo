package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateInvite(ctx context.Context, invite models.Invite) (models.InviteID, error) {
	inviteID, err := uc.ChatServersStorage.CreateInvite(ctx, invite)
	if err != nil {
		return 0, err
	}

	return inviteID, nil
}
