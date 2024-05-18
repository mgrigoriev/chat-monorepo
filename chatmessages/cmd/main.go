package main

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/repository/chatmessages_storage"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/server"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases"
	"log"
)

const grpcPort = "9090"
const httpPort = "8080"
const swaggerPort = "8888"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	repo := chatmessages_storage.New()
	uc := usecases.NewUsecase(usecases.Deps{Repo: repo})

	serverCfg := server.Config{GrpcPort: grpcPort, HttpPort: httpPort, SwaggerPort: swaggerPort}
	serverDeps := server.Deps{Usecase: uc}
	srv, err := server.NewServer(ctx, serverCfg, serverDeps)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	srv.Start(ctx)
}
