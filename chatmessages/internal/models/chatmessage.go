package models

type ChatMessage struct {
	ID            ChatMessageID
	UserID        UserID
	UserName      string
	RecipientType RecipientType // 1: user, 2: server
	RecipientID   UserID
	Content       string
}
