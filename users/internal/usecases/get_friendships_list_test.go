package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_GetFriendshipsList(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx    context.Context
		userID models.UserID
	}

	var friendships = []models.Friendship{
		{
			ID:         1,
			FollowerID: 2,
			FollowedID: 3,
			Status:     "accepted",
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.Friendship
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:    ctx, // dummy
				userID: models.UserID(1),
			},
			want:    &friendships,
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("GetFriendshipsByUserID", ctx, models.UserID(1)).
					Return(&friendships, nil).
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
			got, err := uc.GetFriendshipsByUserID(tt.args.ctx, tt.args.userID)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.GetFriendshipsByUserID() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
