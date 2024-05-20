package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"testing"
)

func Test_usecase_DeleteFriendship(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx          context.Context
		friendshipID models.FriendshipID
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:          ctx, // dummy
				friendshipID: models.FriendshipID(1),
			},
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("DeleteFriendship", ctx, models.FriendshipID(1)).
					Return(nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:          ctx, // dummy
				friendshipID: models.FriendshipID(1),
			},
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.UsersStorage.On("DeleteFriendship", ctx, models.FriendshipID(1)).
					Return(models.ErrDoesNotExist).
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
			err := uc.DeleteFriendship(tt.args.ctx, tt.args.friendshipID)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.DeleteFriendship() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}
