package users_storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	pkgerrors "github.com/mgrigoriev/chat-monorepo/chatservers/pkg/errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

func (r *UsersStorage) GetUsersByTerm(ctx context.Context, term string) (*[]models.User, error) {
	const api = "users_storage.GetUsersByTerm"

	query := squirrel.Select("id", "name", "email", "avatar_photo_url").
		From(usersTable).
		Where("name LIKE ?", fmt.Sprint("%", term, "%")).
		PlaceholderFormat(squirrel.Dollar)

	rows := make([]userRow, 0)
	if err := r.driver.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return nil, pkgerrors.Wrap(api, err)
	}

	users := make([]models.User, 0, len(rows))
	for _, row := range rows {
		users = append(users, models.User{
			ID:             models.UserID(row.ID),
			Name:           row.Name,
			Email:          row.Email,
			AvatarPhotoURL: row.AvatarPhotoURL,
		})
	}

	return &users, nil
}
