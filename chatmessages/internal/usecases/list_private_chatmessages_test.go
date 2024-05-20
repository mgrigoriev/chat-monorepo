package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_ListPrivateChatMessages(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		ChatMessagesStorage *mocks.ChatMessagesStorage
	}
	type args struct {
		ctx         context.Context
		userID      models.UserID
		otherUserID models.UserID
	}

	var friendships = []models.ChatMessage{
		{
			ID:            1,
			UserID:        1,
			UserName:      "John",
			RecipientType: 1,
			RecipientID:   3,
			Content:       "hello",
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.ChatMessage
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:         ctx, // dummy
				userID:      models.UserID(1),
				otherUserID: models.UserID(2),
			},
			want:    &friendships,
			wantErr: nil,
			on: func(f *fields) {
				f.ChatMessagesStorage.On("GetPrivateChatMessages", ctx, models.UserID(1), models.UserID(2)).
					Return(&friendships, nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative.",
			args: args{
				ctx:         ctx, // dummy
				userID:      models.UserID(1),
				otherUserID: models.UserID(2),
			},
			want:    nil,
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.ChatMessagesStorage.On("GetPrivateChatMessages", ctx, models.UserID(1), models.UserID(2)).
					Return(nil, models.ErrDoesNotExist).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			f := &fields{
				ChatMessagesStorage: mocks.NewChatMessagesStorage(t),
			}
			uc := &Usecase{
				Deps: Deps{
					ChatMessagesStorage: f.ChatMessagesStorage,
				},
			}
			if tt.on != nil {
				tt.on(f)
			}

			// act
			got, err := uc.ListPrivateChatMessages(tt.args.ctx, tt.args.userID, tt.args.otherUserID)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.ListPrivateChatMessages() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
