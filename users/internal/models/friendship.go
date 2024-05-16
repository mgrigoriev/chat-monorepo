package models

type Friendship struct {
	ID         FriendshipID
	FollowerID UserID
	FollowedID UserID
	Status     string
}
