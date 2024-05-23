package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) DeleteParticipant(ctx context.Context, participantID models.ParticipantID) error {
	err := uc.ChatServersStorage.DeleteParticipant(ctx, participantID)
	if err != nil {
		return err
	}

	return nil
}
