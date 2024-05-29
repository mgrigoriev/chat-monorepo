package main

import (
	"context"
	"fmt"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/repository/chatmessages_storage"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/postgres"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/transaction_manager"
	"log"
	"os"
	"time"
)

const grpcPort = "9090"
const httpPort = "8080"
const swaggerPort = "8888"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// repository
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "0.0.0.0"
	}

	DSN := fmt.Sprintf("user=mikhail host=%s port=5432 dbname=chatmessages pool_max_conns=10", dbHost)

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
