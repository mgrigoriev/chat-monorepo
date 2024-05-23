package usecases

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

type UsecaseInterface interface {
	CreateChatServer(ctx context.Context, chatserver models.ChatServer) (models.ChatServerID, error)
	GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error)
	SearchChatServers(ctx context.Context, term string) (*[]models.ChatServer, error)
	GetUserChatServers(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error)
	CreateParticipant(ctx context.Context, participant models.Participant) (models.ParticipantID, error)
	DeleteParticipant(ctx context.Context, participantID models.ParticipantID) error
	CreateInvite(ctx context.Context, invite models.Invite) (models.InviteID, error)
}

type ChatServersStorage interface {
	CreateChatServer(ctx context.Context, chatServer models.ChatServer) (models.ChatServerID, error)
	CreateParticipant(ctx context.Context, participant models.Participant) (models.ParticipantID, error)
	CreateInvite(ctx context.Context, invite models.Invite) (models.InviteID, error)
	DeleteParticipant(ctx context.Context, participantID models.ParticipantID) error
	GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error)
	GetChatServersByUserID(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error)
	GetChatServersByTerm(ctx context.Context, term string) (*[]models.ChatServer, error)
}

type TransactionManager interface {
	RunReadCommitted(ctx context.Context, accessMode pgx.TxAccessMode, f func(ctx context.Context) error) error
}

type Deps struct {
	ChatServersStorage
	TransactionManager
}

type Usecase struct {
	Deps
}

func NewUsecase(d Deps) UsecaseInterface {
	return &Usecase{
		Deps: d,
	}
}
