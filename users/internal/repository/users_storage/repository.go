package users_storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	uc "github.com/mgrigoriev/chat-monorepo/users/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/transaction_manager"
)

const (
	usersTable       = "users"
	friendshipsTable = "friendships"
)

// Check that we implement contract for usecase
var (
	_ uc.UsersStorage = (*UsersStorage)(nil)
)

type Connection interface {
	Execx(ctx context.Context, sqlizer postgres.Sqlizer) (pgconn.CommandTag, error)
}

type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) transaction_manager.QueryEngine
}

type UsersStorage struct {
	driver QueryEngineProvider
}

func New(driver QueryEngineProvider) *UsersStorage {
	return &UsersStorage{
		driver: driver,
	}
}
