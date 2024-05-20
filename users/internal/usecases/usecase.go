package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
)

type UsecaseInterface interface {
	CreateUser(ctx context.Context, user models.User) (models.UserID, error)
	GetUserByID(ctx context.Context, id models.UserID) (*models.User, error)
	UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error)
	SearchUsers(ctx context.Context, term string) (*[]models.User, error)
	Login(ctx context.Context, login string, password string) (models.AuthToken, error)
	Auth(ctx context.Context, token models.AuthToken) (*models.User, error)
	CreateFriendship(ctx context.Context, followerID models.UserID, followedID models.UserID) (models.FriendshipID, error)
	GetFriendshipList(ctx context.Context, userID models.UserID) (*[]models.Friendship, error)
	AcceptFriendship(ctx context.Context, friendshipID models.FriendshipID) error
	DeclineFriendship(ctx context.Context, friendshipID models.FriendshipID) error
	DeleteFriendship(ctx context.Context, friendshipID models.FriendshipID) error
}

//go:generate mockery --name=UsersStorage --filename=users_storage_mock.go --disable-version-string
type UsersStorage interface {
	GetUserByToken(ctx context.Context, token models.AuthToken) (*models.User, error)
	GetUserByID(ctx context.Context, id models.UserID) (*models.User, error)
	GetUserByLoginAndPassword(ctx context.Context, login string, password string) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.UserID, error)
	UpdateUser(ctx context.Context, id models.UserID, user models.User) (*models.User, error)
	GetUsersByNameSubstring(ctx context.Context, nameSubstring string) (*[]models.User, error)
	GetFriendshipsByUserID(ctx context.Context, userID models.UserID) (*[]models.Friendship, error)
	CreateFriendship(ctx context.Context, followerID models.UserID, followedID models.UserID) (models.FriendshipID, error)
	UpdateFriendshipStatus(ctx context.Context, friendshipID models.FriendshipID, status string) error
	DeleteFriendship(ctx context.Context, friendshipID models.FriendshipID) error
}

type Deps struct {
	UsersStorage
}

type Usecase struct {
	Deps
}

func NewUsecase(d Deps) UsecaseInterface {
	return &Usecase{
		Deps: d,
	}
}