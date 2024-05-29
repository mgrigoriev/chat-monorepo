package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/models"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_ListServerChatMessages(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		ChatMessagesStorage *mocks.ChatMessagesStorage
	}
	type args struct {
		ctx      context.Context
		serverID models.ChatServerID
	}

	var chatMessages = []models.ChatMessage{
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
				ctx:      ctx, // dummy
				serverID: models.ChatServerID(1),
			},
			want:    &chatMessages,
			wantErr: nil,
			on: func(f *fields) {
				f.ChatMessagesStorage.On("GetServerChatMessages", ctx, models.ChatServerID(1)).
					Return(&chatMessages, nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative.",
			args: args{
				ctx:      ctx, // dummy
				serverID: models.ChatServerID(1),
			},
			want:    nil,
			wantErr: models.ErrDoesNotExist,
			on: func(f *fields) {
				f.ChatMessagesStorage.On("GetServerChatMessages", ctx, models.ChatServerID(1)).
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
			got, err := uc.ListServerChatMessages(tt.args.ctx, tt.args.serverID)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.ListServerChatMessages() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
