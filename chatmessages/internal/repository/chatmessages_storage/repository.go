package chatmessages_storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	uc "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/transaction_manager"
)

const (
	chatmessagesTable = "chatmessages"
)

// Check that we implement contract for usecase
var (
	_ uc.ChatMessagesStorage = (*ChatMessagesStorage)(nil)
)

type Connection interface {
	Execx(ctx context.Context, sqlizer postgres.Sqlizer) (pgconn.CommandTag, error)
}

type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) transaction_manager.QueryEngine
}

type ChatMessagesStorage struct {
	driver QueryEngineProvider
}

func New(driver QueryEngineProvider) *ChatMessagesStorage {
	return &ChatMessagesStorage{
		driver: driver,
	}
}
