package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_SearchUsers(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx  context.Context
		term string
	}

	var users = []models.User{
		{
			ID:             1,
			Name:           "Test User",
			Email:          "test@test.com",
			AvatarPhotoURL: "https://test.com/test.jpg",
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.User
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:  ctx, // dummy
				term: "john",
			},
			want:    &users,
			wantErr: nil,
			on: func(f *fields) {
				f.UsersStorage.On("GetUsersByNameSubstring", ctx, "john").
					Return(&users, nil).
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
			got, err := uc.SearchUsers(tt.args.ctx, tt.args.term)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.SearchUsers() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
