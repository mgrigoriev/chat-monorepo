package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

func (uc *Usecase) CreateParticipant(ctx context.Context, participant models.Participant) (models.ParticipantID, error) {
	participantID, err := uc.ChatServersStorage.CreateParticipant(ctx, participant)
	if err != nil {
		return 0, err
	}

	return participantID, nil
}
