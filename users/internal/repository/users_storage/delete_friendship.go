package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) DeleteFriendship(friendshipID models.FriendshipID) error {
	_ = friendshipID

	return nil
}
