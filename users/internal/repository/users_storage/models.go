package users_storage

import "github.com/mgrigoriev/chat-monorepo/users/internal/models"

type userRow struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	Email          string `db:"email"`
	Password       string `db:"password"` // TODO: hash password
	AvatarPhotoURL string `db:"avatar_photo_url"`
}

type friendshipRow struct {
	ID         int64  `db:"id"`
	FollowerID int64  `db:"follower_id"`
	FollowedID int64  `db:"followed_id"`
	Status     string `db:"status"`
}

func (r *userRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":    r.ID,
		"name":             r.Name,
		"email":            r.Email,
		"password":         r.Password,
		"avatar_photo_url": r.AvatarPhotoURL,
	}
}

func (r *friendshipRow) ValuesMap() map[string]any {
	return map[string]any{
		//"id":           r.ID,
		"follower_id": r.FollowerID,
		"followed_id": r.FollowedID,
		"status":      r.Status,
	}
}

func (r *userRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func (r *friendshipRow) Values(columns ...string) []any {
	values := make([]any, 0, len(columns))
	m := r.ValuesMap()

	for i := range columns {
		values = append(values, m[columns[i]])
	}

	return values
}

func newUserRowFromModel(user *models.User) (*userRow, error) {
	return &userRow{
		ID:             int64(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}, nil
}

func newFriendshipRowFromModel(friendship *models.Friendship) (*friendshipRow, error) {
	return &friendshipRow{
		ID:         int64(friendship.ID),
		FollowerID: int64(friendship.FollowerID),
		FollowedID: int64(friendship.FollowedID),
		Status:     friendship.Status,
	}, nil
}
