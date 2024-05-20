package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_Login(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}

	var user = models.User{
		ID:             3,
		Name:           "Test",
		Email:          "test@test.com",
		AvatarPhotoURL: "https://test.com/1.jpg",
	}

	tests := []struct {
		name    string
		args    args
		want    models.AuthToken
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:      ctx, // dummy
				login:    "test@example.com",
				password: "correct-password",
			},
			want:    models.AuthToken("valid-token"),
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByLoginAndPassword", ctx, "test@example.com", "correct-password").
					Return(&user, nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:      ctx, // dummy
				login:    "test@example.com",
				password: "wrong-password",
			},
			want:    models.AuthToken(""),
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByLoginAndPassword", ctx, "test@example.com", "wrong-password").
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
			got, err := uc.Login(tt.args.ctx, tt.args.login, tt.args.password)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.Login() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
