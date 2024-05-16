package models

type User struct {
	ID             UserID
	Name           string
	Email          string
	Password       string
	AvatarPhotoURL string
}
