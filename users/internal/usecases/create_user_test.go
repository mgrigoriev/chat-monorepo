package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_CreateUser(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx  context.Context
		user models.User
	}
	var user = models.User{
		Name:           "Test User",
		Email:          "test@test.com",
		Password:       "qwerty",
		AvatarPhotoURL: "https://test.com/test.jpg",
	}

	tests := []struct {
		name    string
		args    args
		want    models.UserID
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:  ctx, // dummy
				user: user,
			},
			want:    models.UserID(1),
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("CreateUser", ctx, user).
					Return(models.UserID(1), nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:  ctx, // dummy
				user: user,
			},
			want:    models.UserID(0),
			wantErr: models.ErrAlreadyExists,
			on: func(f *fields) {
				f.UsersStorage.On("CreateUser", ctx, user).
					Return(models.UserID(0), models.ErrAlreadyExists).
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
			got, err := uc.CreateUser(tt.args.ctx, tt.args.user)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.CreateUser() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
