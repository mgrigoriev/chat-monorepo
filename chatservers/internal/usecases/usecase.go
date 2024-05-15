package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
)

type UsecaseInterface interface {
	CreateChatServer(ctx context.Context, chatserver models.ChatServer) (models.ChatServerID, error)
	GetChatServerByID(ctx context.Context, id models.ChatServerID) (*models.ChatServer, error)
	SearchChatServers(ctx context.Context, term string) (*[]models.ChatServer, error)
	GetUserChatServers(ctx context.Context, userID models.UserID) (*[]models.ChatServer, error)
	CreateParticipant(ctx context.Context, chatServerID models.ChatServerID, userID models.UserID) (models.ParticipantID, error)
	DeleteParticipant(ctx context.Context, participantID models.ParticipantID) error
	CreateInvite(ctx context.Context, chatServerID models.ChatServerID, userID models.UserID) (models.InviteID, error)
}

type ChatServersRepoInterface interface{}

type Deps struct {
	Repo ChatServersRepoInterface
}

type Usecase struct {
	Deps
}

func NewUsecase(d Deps) UsecaseInterface {
	return &Usecase{
		Deps: d,
	}
}
