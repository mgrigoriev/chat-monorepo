package usecases

import (
	"context"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_usecase_CreateUser(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
		//date = time.Now()
	)
	type fields struct {
		UsersStorage *mocks.UsersStorage
	}
	type args struct {
		ctx  context.Context
		user models.User
	}

	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool

		on     func(*fields)
		assert func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx: ctx, // dummy
				user: models.User{
					Name:           "Test User",
					Email:          "test@test.com",
					Password:       "qwerty",
					AvatarPhotoURL: "https://test.com/test.jpg",
				},
			},
			want: &models.User{
				ID:             1,
				Name:           "Test User",
				Email:          "test@test.com",
				Password:       "qwerty",
				AvatarPhotoURL: "https://test.com/test.jpg",
			},
			wantErr: false,

			on: func(f *fields) {
				f.UsersStorage.On("CreateUser", ctx, mock.MatchedBy(func(user *models.User) bool {
					return user == &models.User{
						Name:           "Test User",
						Email:          "test@test.com",
						Password:       "qwerty",
						AvatarPhotoURL: "https://test.com/test.jpg",
					}
				})).Return(nil)
			},
			assert: func(t *testing.T, f *fields) {
				f.UsersStorage.AssertNumberOfCalls(t, "CreateUser", 1)
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
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil { // зануляем так как не можем проверить
				got.ID = models.OrderID{}
			}
			assert.Equal(t, tt.want, got)

			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}
