package main

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/repository/chatmessages_storage"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/server"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/pkg/transaction_manager"
	"log"
	"time"
)

const grpcPort = "9090"
const httpPort = "8080"
const swaggerPort = "8888"

const DSN = "user=mikhail host=host.docker.internal port=5432 dbname=chatmessages pool_max_conns=10"

// const DSN = "user=mikhail host=0.0.0.0 port=5432 dbname=chatmessages pool_max_conns=10"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
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
	storage := chatmessages_storage.New(txManager)

	// services

	// ...

	// usecases

	uc := usecases.NewUsecase(usecases.Deps{
		ChatMessagesStorage: storage,
		TransactionManager:  txManager,
	})

	serverCfg := server.Config{GrpcPort: grpcPort, HttpPort: httpPort, SwaggerPort: swaggerPort}
	serverDeps := server.Deps{Usecase: uc}
	srv, err := server.NewServer(ctx, serverCfg, serverDeps)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	srv.Start(ctx)
}
