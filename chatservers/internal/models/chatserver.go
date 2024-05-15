package models

type ChatServer struct {
	ID     ChatServerID
	UserID UserID // Creator ID
	Name   string // Chat server name
}
