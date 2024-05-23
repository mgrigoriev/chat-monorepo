package main

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/repository/chatservers_storage"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/transaction_manager"
	"log"
	"time"
)

const DSN = "user=mikhail host=0.0.0.0 port=5432 dbname=chatservers pool_max_conns=10"
const port = "8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// repository
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

	s := server.New(ctx, serverCfg, serverDeps)
	s.Start()
}
