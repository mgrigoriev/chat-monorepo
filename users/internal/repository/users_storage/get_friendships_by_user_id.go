package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

func (r *UsersStorage) GetFriendshipsByUserID(userID models.UserID) (*[]models.Friendship, error) {
	_ = userID

	return &[]models.Friendship{
		{
			ID:         10,
			FollowerID: 1,
			FollowedID: 2,
			Status:     "approved",
		},
	}, nil
}
