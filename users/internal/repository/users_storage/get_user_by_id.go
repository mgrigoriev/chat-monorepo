package users_storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) GetUserByID(ctx context.Context, id models.UserID) (*models.User, error) {
	const api = "users_storage.GetUserByID"

	query := squirrel.Select("id", "name", "email", "avatar_photo_url").
		From(usersTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	var row userRow
	err := r.driver.GetQueryEngine(ctx).Getx(ctx, &row, query)
	if err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	user := models.User{
		ID:             models.UserID(row.ID),
		Name:           row.Name,
		Email:          row.Email,
		AvatarPhotoURL: row.AvatarPhotoURL,
	}

	return &user, nil
}
