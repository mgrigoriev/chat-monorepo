package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetUserByToken(token models.AuthToken) (*models.User, error) {
	_ = token

	return &models.User{
		ID:             1,
		Name:           "Test 1",
		Email:          "test1@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}

func (r *Repo) GetUserByID(id models.UserID) (*models.User, error) {
	return &models.User{
		ID:             id,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}

func (r *Repo) GetUserByLoginAndPassword(login string, password string) (*models.User, error) {
	_ = login
	_ = password

	return &models.User{
		ID:             3,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}, nil
}

func (r *Repo) CreateUser(user models.User) (models.UserID, error) {
	_ = user

	return models.UserID(1), nil
}

func (r *Repo) UpdateUser(id models.UserID, user models.User) (*models.User, error) {
	return &models.User{
		ID:             id,
		Name:           "Updated User",
		Email:          user.Email,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}, nil
}

func (r *Repo) GetUsersByNameSubstring(nameSubstring string) (*[]models.User, error) {
	_ = nameSubstring

	return &[]models.User{
		{
			ID:             1,
			Name:           "Test",
			Email:          "test@test.com",
			AvatarPhotoURL: "https://test.com/1.jpg",
		},
	}, nil
}

func (r *Repo) GetFriendshipsByUserID(userID models.UserID) (*[]models.Friendship, error) {
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

func (r *Repo) CreateFriendship(followerID models.UserID, followedID models.UserID) (models.FriendshipID, error) {
	_ = followerID
	_ = followedID

	return models.FriendshipID(1), nil
}

func (r *Repo) UpdateFriendshipStatus(friendshipID models.FriendshipID, status string) error {
	_ = friendshipID
	_ = status

	return nil
}

func (r *Repo) DeleteFriendship(friendshipID models.FriendshipID) error {
	_ = friendshipID

	return nil
}
