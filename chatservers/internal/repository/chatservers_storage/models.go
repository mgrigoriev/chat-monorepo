package chatservers_storage

import (
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

type chatServerRow struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
}

type participantRow struct {
	ID           int64 `db:"id"`
	UserID       int64 `db:"user_id"`
	ChatServerID int64 `db:"chatserver_id"`
}

type inviteRow struct {
	ID           int64 `db:"id"`
	UserID       int64 `db:"user_id"`
	ChatServerID int64 `db:"chatserver_id"`
}

func (r *chatServerRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":    r.ID,
		"user_id": r.UserID,
		"name":    r.Name,
	}
}

func (r *participantRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":           r.ID,
		"user_id":       r.UserID,
		"chatserver_id": r.ChatServerID,
	}
}

func (r *inviteRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":           r.ID,
		"user_id":       r.UserID,
		"chatserver_id": r.ChatServerID,
	}
}

func (r *chatServerRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func (r *participantRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func (r *inviteRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func newChatServerRowFromModel(chatServer *models.ChatServer) (*chatServerRow, error) {
	return &chatServerRow{
		ID:     int64(chatServer.ID),
		UserID: int64(chatServer.UserID),
		Name:   chatServer.Name,
	}, nil
}

func newParticipantRowFromModel(participant *models.Participant) (*participantRow, error) {
	return &participantRow{
		ID:           int64(participant.ID),
		UserID:       int64(participant.UserID),
		ChatServerID: int64(participant.ChatServerID),
	}, nil
}

func newInviteRowFromModel(invite *models.Invite) (*inviteRow, error) {
	return &inviteRow{
		ID:           int64(invite.ID),
		UserID:       int64(invite.UserID),
		ChatServerID: int64(invite.ChatServerID),
	}, nil
}
