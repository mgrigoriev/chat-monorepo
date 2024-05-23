package chatservers_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatservers/pkg/errors"
)

func (r *ChatServersStorage) DeleteParticipant(ctx context.Context, participantID models.ParticipantID) error {
	const api = "chatservers_storage.DeleteParticipant"

	query := squirrel.Delete(participantsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": participantID})

	if _, err := r.driver.GetQueryEngine(ctx).Execx(ctx, query); err != nil {
		return pkgerrors.Wrap(api, err)
	}

	return nil
}
