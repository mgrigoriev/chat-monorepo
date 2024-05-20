package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_GetUserByID(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx context.Context
		id  models.UserID
	}

	var user = models.User{
		ID:             1,
		Name:           "Test User",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/test.jpg",
	}

	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx: ctx, // dummy
				id:  models.UserID(1),
			},
			want:    &user,
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByID", ctx, models.UserID(1)).
					Return(&user, nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx: ctx, // dummy
				id:  models.UserID(1),
			},
			want:    nil,
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByID", ctx, models.UserID(1)).
					Return(nil, models.ErrDoesNotExist).
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
			got, err := uc.GetUserByID(tt.args.ctx, tt.args.id)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.GetUserByID() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
