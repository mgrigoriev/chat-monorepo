package main

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/repository/users_storage"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases"
)

const port = "8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := users_storage.New()
	uc := usecases.NewUsecase(usecases.Deps{Repo: repo})

	serverCfg := server.Config{Port: port}
	serverDeps := server.Deps{Usecase: uc}

	s := server.New(ctx, serverCfg, serverDeps)
	s.Start()
}
