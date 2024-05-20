package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_CreateFriendship(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx        context.Context
		followerID models.UserID
		followedID models.UserID
	}

	tests := []struct {
		name    string
		args    args
		want    models.FriendshipID
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:        ctx, // dummy
				followerID: models.UserID(1),
				followedID: models.UserID(2),
			},
			want:    models.FriendshipID(123),
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("CreateFriendship", ctx, models.UserID(1), models.UserID(2)).
					Return(models.FriendshipID(123), nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:        ctx, // dummy
				followerID: models.UserID(1),
				followedID: models.UserID(2),
			},
			want:    models.FriendshipID(0),
			wantErr: models.ErrAlreadyExists,
			on: func(f *fields) {
				f.UsersStorage.On("CreateFriendship", ctx, models.UserID(1), models.UserID(2)).
					Return(models.FriendshipID(0), models.ErrAlreadyExists).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			f := &fields{
				UsersStorage: mocks.NewUsersStorage(t),
			}
			uc := &Usecase{
				Deps: Deps{
					UsersStorage: f.UsersStorage,
				},
			}
			if tt.on != nil {
				tt.on(f)
			}

			// act
			got, err := uc.CreateFriendship(tt.args.ctx, tt.args.followerID, tt.args.followedID)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.CreateFriendship() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
