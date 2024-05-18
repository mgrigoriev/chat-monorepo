package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
)

type UsecaseInterface interface {
	SaveChatMessage(ctx context.Context, chatmessage models.ChatMessage) (models.ChatMessageID, error)
	ListPrivateChatMessages(ctx context.Context, userID models.UserID, otherUserID models.UserID) (*[]models.ChatMessage, error)
	ListServerChatMessages(ctx context.Context, serverID models.ChatServerID) (*[]models.ChatMessage, error)
}

type ChatMessagesRepoInterface interface{}

type Deps struct {
	Repo ChatMessagesRepoInterface
}

type Usecase struct {
	Deps
}

func NewUsecase(d Deps) UsecaseInterface {
	return &Usecase{
		Deps: d,
	}
}

func (uc *Usecase) SaveChatMessage(ctx context.Context, chatmessage models.ChatMessage) (models.ChatMessageID, error) {
	// return 0, models.ErrNotImplemented

	// TODO: Save to repo

	return 1, nil
}
