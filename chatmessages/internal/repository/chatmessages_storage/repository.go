package chatmessages_storage

import uc "github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases"

// Check that we implement contract for usecase
var (
	_ uc.ChatMessagesStorage = (*ChatMessagesStorage)(nil)
)

type ChatMessagesStorage struct{}

func New() *ChatMessagesStorage {
	return &ChatMessagesStorage{}
}
