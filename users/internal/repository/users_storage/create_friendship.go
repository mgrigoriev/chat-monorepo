package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) CreateFriendship(followerID models.UserID, followedID models.UserID) (models.FriendshipID, error) {
	_ = followerID
	_ = followedID

	return models.FriendshipID(1), nil
}
