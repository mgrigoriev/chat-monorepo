package main

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/users/internal/repository/users_storage"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/logger"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/transaction_manager"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const port = "8080"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	logger.SetLevel(zapcore.DebugLevel)
	logger.Info(ctx, "start app init")

	// repository
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "0.0.0.0"
	}

	DSN := fmt.Sprintf("user=mikhail host=%s port=5432 dbname=users pool_max_conns=10", dbHost)

	pool, err := postgres.NewConnectionPool(ctx, DSN,
		postgres.WithMaxConnIdleTime(5*time.Minute),
		postgres.WithMaxConnLifeTime(time.Hour),
		postgres.WithMaxConnectionsCount(10),
		postgres.WithMinConnectionsCount(5),
	)
	if err != nil {
		logger.Fatal(ctx, err)
	}

	txManager := transaction_manager.New(pool)
	storage := users_storage.New(txManager)

	// usecases
	uc := usecases.NewUsecase(usecases.Deps{
		UsersStorage:       storage,
		TransactionManager: txManager,
	})

	serverCfg := server.Config{Port: port}
	serverDeps := server.Deps{Usecase: uc}

	s := server.New(serverCfg, serverDeps)
	if err := s.Start(ctx); err != nil {
		logger.Fatal(ctx, err)
	}
}
