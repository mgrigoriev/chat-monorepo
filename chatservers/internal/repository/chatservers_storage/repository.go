package chatservers_storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	uc "github.com/mgrigoriev/chat-monorepo/chatservers/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/transaction_manager"
)

const (
	chatserversTable  = "chatservers"
	participantsTable = "participants"
	invitesTable      = "invites"
)

// Check that we implement contract for usecase
var (
	_ uc.ChatServersStorage = (*ChatServersStorage)(nil)
)

type Connection interface {
	Execx(ctx context.Context, sqlizer postgres.Sqlizer) (pgconn.CommandTag, error)
}

type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) transaction_manager.QueryEngine
}

type ChatServersStorage struct {
	driver QueryEngineProvider
}

func New(driver QueryEngineProvider) *ChatServersStorage {
	return &ChatServersStorage{
		driver: driver,
	}
}
