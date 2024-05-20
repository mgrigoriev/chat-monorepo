package usecases

import (
	"context"
	"errors"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/models"
	"github.com/mgrigoriev/chat-monorepo/chatmesages/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_usecase_SaveChatMessage(t *testing.T) {
	// prepare
	var (
		ctx = context.Background() // dummy
	)
	type fields struct {
		ChatmessagesStorage *mocks.ChatMessagesStorage
	}
	type args struct {
		ctx         context.Context
		chatMessage models.ChatMessage
	}
	var chatMessage = models.ChatMessage{
		ID:            1,
		UserID:        2,
		UserName:      "",
		RecipientType: 1,
		RecipientID:   3,
		Content:       "hello",
	}

	tests := []struct {
		name    string
		args    args
		want    models.ChatMessageID
		wantErr error
		on      func(*fields)
		assert  func(*testing.T, *fields)
	}{
		{
			name: "Test 1. Positive.",
			args: args{
				ctx:         ctx, // dummy
				chatMessage: chatMessage,
			},
			want:    models.ChatMessageID(1),
			wantErr: nil,
			on: func(f *fields) {
				f.ChatmessagesStorage.On("CreateChatMessage", ctx, chatMessage).
					Return(models.ChatMessageID(1), nil).
					Once()
			},
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx:         ctx, // dummy
				chatMessage: chatMessage,
			},
			want:    models.ChatMessageID(0),
			wantErr: models.ErrAlreadyExists,
			on: func(f *fields) {
				f.ChatmessagesStorage.On("CreateChatMessage", ctx, chatMessage).
					Return(models.ChatMessageID(0), models.ErrAlreadyExists).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			f := &fields{
				ChatmessagesStorage: mocks.NewChatMessagesStorage(t),
			}
			uc := &Usecase{
				Deps: Deps{
					ChatMessagesStorage: f.ChatmessagesStorage,
				},
			}
			if tt.on != nil {
				tt.on(f)
			}

			// act
			got, err := uc.CreateChatMessage(tt.args.ctx, tt.args.chatMessage)

			// assert
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("usecase.SaveChatMessage() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
