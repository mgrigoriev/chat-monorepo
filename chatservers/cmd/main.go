package main

import (
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
