package main

import "github.com/mgrigoriev/chat-monorepo/users/internal/server"

func main() {
	s := server.NewServer()
	s.Start()
}
