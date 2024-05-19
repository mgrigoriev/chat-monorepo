package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) UpdateFriendshipStatus(friendshipID models.FriendshipID, status string) error {
	_ = friendshipID
	_ = status

	return nil
}
