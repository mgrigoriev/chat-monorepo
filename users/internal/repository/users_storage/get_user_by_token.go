package users_storage

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/users/pkg/errors"
)

func (r *UsersStorage) GetUserByToken(ctx context.Context, token models.AuthToken) (*models.User, error) {
	const api = "users_storage.GetUserByToken"

	// Hardcoded token for now
	if token != "valid-token" {
		return nil, pkgerrors.Wrap(api, errors.New("invalid token"))
	}

	// Hardcode user ID
	query := squirrel.Select("id", "name", "email", "avatar_photo_url").
		From(usersTable).
		Where(squirrel.Eq{"id": 2}).
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
