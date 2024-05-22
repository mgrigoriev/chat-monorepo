package usecases

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

type UsecaseInterface interface {
	SaveChatMessage(ctx context.Context, chatmessage models.ChatMessage) (models.ChatMessageID, error)
	ListPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error)
	ListServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error)
}

//go:generate mockery --name=ChatMessagesStorage --filename=chatmessages_storage_mock.go --disable-version-string
type ChatMessagesStorage interface {
	CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) (models.ChatMessageID, error)
	GetPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error)
	GetServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error)
}

type TransactionManager interface {
	RunReadCommitted(ctx context.Context, accessMode pgx.TxAccessMode, f func(ctx context.Context) error) error
}

type Deps struct {
	ChatMessagesStorage
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
