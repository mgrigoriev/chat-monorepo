package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateParticipant(ctx context.Context, chatServerID models.ChatServerID, userID models.UserID) (models.ParticipantID, error) {
	//return 0, models.ErrNotImplemented
	return 10, nil
}
