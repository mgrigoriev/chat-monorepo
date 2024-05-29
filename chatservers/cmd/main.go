package main

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/repository/chatservers_storage"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/transaction_manager"
	"log"
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

	// repository
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "0.0.0.0"
	}

	DSN := fmt.Sprintf("user=mikhail host=%s port=5432 dbname=chatservers pool_max_conns=10", dbHost)

	pool, err := postgres.NewConnectionPool(ctx, DSN,
		postgres.WithMaxConnIdleTime(5*time.Minute),
		postgres.WithMaxConnLifeTime(time.Hour),
		postgres.WithMaxConnectionsCount(10),
		postgres.WithMinConnectionsCount(5),
	)
	if err != nil {
		log.Fatal(err)
	}

	txManager := transaction_manager.New(pool)
	storage := chatservers_storage.New(txManager)

	// usecases
	uc := usecases.NewUsecase(usecases.Deps{
		ChatServersStorage: storage,
		TransactionManager: txManager,
	})

	serverCfg := server.Config{Port: port}
	serverDeps := server.Deps{Usecase: uc}

	s := server.New(serverCfg, serverDeps)
	if err := s.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
