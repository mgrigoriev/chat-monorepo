package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_Auth(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx   context.Context
		token models.AuthToken
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
				ctx:   ctx, // dummy
				token: models.AuthToken("valid"),
			},
			want:    &user,
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByToken", ctx, models.AuthToken("valid")).
					Return(&user, nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:   ctx, // dummy
				token: models.AuthToken("invalid"),
			},
			want:    nil,
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.UsersStorage.On("GetUserByToken", ctx, models.AuthToken("invalid")).
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
			got, err := uc.Auth(tt.args.ctx, tt.args.token)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.Auth() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
