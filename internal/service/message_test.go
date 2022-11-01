package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/messagetemplate"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMessage_GenerateID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	userID := int64(1)
	ts := time.Now().UnixMilli()
	externalID := int64(2)
	expectedID := fmt.Sprintf("NS-%03d-%03d-%019d-%d-%019d",
		channel.Telegram,
		messagetemplate.Receipt,
		userID,
		ts,
		externalID)

	id := services.Message.GenerateID(channel.Telegram, messagetemplate.Receipt, userID, ts, externalID)
	assert.Equal(t, expectedID, id)

	userID = int64(1)
	ts = time.Now().Unix()
	externalID = int64(2)
	expectedID = fmt.Sprintf("NS-%03d-%03d-%019d-%d000-%019d",
		channel.Telegram,
		messagetemplate.Receipt,
		userID,
		ts,
		externalID)

	id = services.Message.GenerateID(channel.Telegram, messagetemplate.Receipt, userID, ts, externalID)
	assert.Equal(t, expectedID, id)
}

func TestMessage_ValidateID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name               string
		input              string
		userID             int64
		expectedError      error
		expectedQueryError error
	}{
		{
			name:   "ok",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
		},
		{
			name:          "empty id",
			input:         "",
			expectedError: InvalidMessageIDErr,
		},
		{
			name:          "invalid id",
			input:         "NS-001-001-0000000001234567890-1166115824000",
			expectedError: InvalidMessageIDErr,
		},
		{
			name:          "invalid channel",
			input:         "NS-000-001-0000000001234567890-1666115824000-0000000001234567897",
			expectedError: InvalidChannelErr,
		},
		{
			name:          "invalid message template",
			input:         "NS-001-000-0000000001234567890-1666115824000-0000000001234567897",
			expectedError: InvalidMessageTemplateErr,
		},
		{
			name:          "user not exists",
			input:         "NS-001-001-0000000001234567891-1166115824000-0000000001234567897",
			userID:        1234567891,
			expectedError: InvalidUserErr,
		},
		{
			name:          "invalid timestamp",
			input:         "NS-001-001-0000000001234567890-1166115824000-0000000001234567897",
			userID:        1234567890,
			expectedError: InvalidTimestamp,
		},
		{
			name:               "database error",
			input:              "NS-001-001-0000000001234567890-1166115824000-0000000001234567897",
			userID:             1234567890,
			expectedQueryError: errors.New("database error :("),
		},
	}

	ctx := context.Background()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.userID != 0 {
				mocked.RepositoryUser.EXPECT().Exists(ctx, tc.userID).Return(!errors.Is(tc.expectedError, InvalidUserErr), tc.expectedQueryError)
			}

			err := services.Message.ValidateID(ctx, tc.input)

			if tc.expectedQueryError != nil {
				assert.Equal(t, tc.expectedQueryError, err)
			} else {
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}

func TestMessage_parseID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name     string
		input    string
		expected *entity.Message
	}{
		{
			name:  "ok",
			input: "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			expected: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
		},
	}

	ctx := context.Background()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != nil {
				mocked.RepositoryUser.EXPECT().Exists(ctx, tc.expected.UserID).Return(true, nil)
			}

			data, err := services.Message.parseID(ctx, tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, data)
		})
	}
}

func TestMessage_GetStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name               string
		input              string
		userID             int64
		expected           *entity.MessageStatus
		expectedError      error
		expectedQueryError error
		expectedIDError    error
	}{
		{
			name:   "ok",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			expected: &entity.MessageStatus{
				MessageID:   "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Status:      entity.MessageStatusSending,
				Description: "",
				CreatedAt:   time.Now(),
			},
		},
		{
			name:               "not found message",
			input:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567891",
			userID:             1234567890,
			expectedQueryError: sql.ErrNoRows,
			expectedError:      MessageNotFoundErr,
		},
		{
			name:               "unknown error",
			input:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567891",
			userID:             1234567890,
			expectedQueryError: fmt.Errorf("unknown error"),
			expectedError:      fmt.Errorf("unknown error"),
		},
		{
			name:            "invalid id timestamp",
			input:           "NS-001-001-0000000001234567890-1166115824000-0000000001234567891",
			userID:          1234567890,
			expectedIDError: InvalidTimestamp,
		},
	}

	ctx := context.Background()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mocked.RepositoryUser.EXPECT().Exists(ctx, tc.userID).Return(true, tc.expectedIDError)

			if tc.expectedIDError == nil {
				mocked.RepositoryMessage.EXPECT().FindLastStatus(ctx, tc.input).Return(tc.expected, tc.expectedQueryError)
			}

			data, err := services.Message.GetStatus(ctx, tc.input)
			assert.Equal(t, tc.expected, data)

			if tc.expectedIDError == nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Equal(t, fmt.Errorf("get status: %w", tc.expectedIDError), err)
			}
		})
	}
}

func TestMessage_Send(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name                        string
		input                       string
		userID                      int64
		message                     *entity.Message
		expectedId                  string
		expectedErrorId             error
		expectedErrorOnDuplicate    error
		expectedErrorOnCreate       error
		expectedErrorOnCreateStatus error
		expectedErrorOnPublish      error
	}{
		{
			name:   "ok",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedId: "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
		},
		{
			name:   "duplicate",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedId: "NS-001-001-0000000001234567890-1666115823000-0000000001234567890",
		},
		{
			name:            "invalid id",
			input:           "",
			expectedErrorId: InvalidMessageIDErr,
		},
		{
			name:   "query error 'CheckForDuplicates'",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedErrorOnDuplicate: fmt.Errorf("database error :("),
		},
		{
			name:   "query error 'Create'",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedErrorOnCreate: fmt.Errorf("database error :("),
		},
		{
			name:   "query error 'CreateStatus'",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedErrorOnCreateStatus: fmt.Errorf("database error :("),
		},
		{
			name:   "query error 'Publish'",
			input:  "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
			userID: 1234567890,
			message: &entity.Message{
				ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
				Channel:         channel.Telegram,
				UserID:          1234567890,
				MessageTemplate: messagetemplate.Receipt,
				Timestamp:       1666115824000,
				ExternalID:      1234567890,
			},
			expectedErrorOnCreateStatus: fmt.Errorf("message broker error :("),
		},
	}

	ctx := context.Background()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var testErr error

			defer func() {
				id, err := services.Message.Send(ctx, tc.input, nil)
				assert.Equal(t, testErr, err)
				assert.Equal(t, tc.expectedId, id)
			}()

			if tc.expectedErrorId != nil {
				testErr = tc.expectedErrorId
				return
			}

			mocked.RepositoryUser.EXPECT().Exists(ctx, tc.userID).Return(true, nil)

			if tc.expectedErrorOnDuplicate != nil {
				mocked.RepositoryMessage.EXPECT().CheckForDuplicates(ctx, tc.message).Return("", tc.expectedErrorOnDuplicate)
				testErr = tc.expectedErrorOnDuplicate
				return
			} else if tc.expectedId != "" && tc.expectedId != tc.input {
				mocked.RepositoryMessage.EXPECT().CheckForDuplicates(ctx, tc.message).Return(tc.expectedId, nil)
				return
			} else {
				mocked.RepositoryMessage.EXPECT().CheckForDuplicates(ctx, tc.message).Return("", sql.ErrNoRows)
				testErr = sql.ErrNoRows
			}

			mocked.RepositoryMessage.EXPECT().Create(ctx, tc.message).Return(tc.expectedErrorOnCreate)
			if tc.expectedErrorOnCreate != nil {
				testErr = tc.expectedErrorOnCreate
				return
			}

			mocked.RepositoryMessage.EXPECT().CreateStatus(ctx, tc.message.ID, entity.MessageStatusNew, "").Return(tc.expectedErrorOnCreateStatus)
			if tc.expectedErrorOnCreateStatus != nil {
				testErr = tc.expectedErrorOnCreateStatus
				return
			}

			mocked.RepositoryMessage.EXPECT().Publish(ctx, services.Message.pubSubKey(tc.message.Channel), tc.message.ID).Return(tc.expectedErrorOnPublish)
			testErr = tc.expectedErrorOnPublish
		})
	}
}

func TestMessage_getChannelNameFromPubSubKey(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name          string
		input         string
		expected      channel.Channel
		expectedError error
	}{
		{
			name:          "ok telegram",
			input:         fmt.Sprintf("ns::%d", channel.Telegram),
			expected:      channel.Telegram,
			expectedError: nil,
		},
		{
			name:          "ok email",
			input:         fmt.Sprintf("ns::%d", channel.Email),
			expected:      channel.Email,
			expectedError: nil,
		},
		{
			name:          "invalid input",
			input:         "123",
			expectedError: InvalidChannelErr,
		},
		{
			name:          "invalid channel",
			input:         "ns::0",
			expectedError: InvalidChannelErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ch, err := services.Message.getChannelNameFromPubSubKey(tc.input)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expected, ch)
		})
	}
}

func TestMessage_makeStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name                           string
		messageID, status, description string
		expectedError                  error
	}{
		{
			name:          "ok",
			messageID:     "123",
			status:        entity.MessageStatusNew,
			description:   "testDescriptionNew",
			expectedError: nil,
		},
		{
			name:          "with error",
			messageID:     "123",
			status:        entity.MessageStatusSent,
			description:   "testDescriptionSent",
			expectedError: errors.New("database error :("),
		},
	}

	ctx := context.Background()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mocked.RepositoryMessage.EXPECT().CreateStatus(ctx, tc.messageID, tc.status, tc.description).Return(tc.expectedError)

			if tc.expectedError != nil {
				mocked.Logger.EXPECT().Error("create message status",
					"messageId", tc.messageID,
					"status", tc.status,
					"description", tc.description,
					"error", tc.expectedError)
			}

			services.Message.makeStatus(ctx, tc.messageID, tc.status, tc.description)
		})
	}
}

func TestMessage_getContentFromTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name            string
		input           *entity.Message
		expectedError   error
		expectedContent string
	}{
		{
			name: "invalid message template type",
			input: &entity.Message{
				Channel:         channel.Telegram,
				MessageTemplate: messagetemplate.MessageTemplate(0),
				Params:          []byte(`{"orderId": 123, "commissionAmount": "1 KZT", "totalAmount": "1001 KZT"}`),
			},
			expectedError: fmt.Errorf("failed to get message template: %w", messagetemplate.TemplateNotFoundErr),
		},
		{
			name: "invalid params",
			input: &entity.Message{
				Channel:         channel.Telegram,
				MessageTemplate: messagetemplate.Receipt,
				Params:          nil,
			},
			expectedError: fmt.Errorf("failed to set params message template: %s", errors.New("unexpected end of JSON input")),
		},
		{
			name: "invalid channel",
			input: &entity.Message{
				Channel:         channel.Channel(0),
				MessageTemplate: messagetemplate.Receipt,
				Params:          []byte(`{"orderId": 123, "commissionAmount": "1 KZT", "totalAmount": "1001 KZT"}`),
			},
			expectedError: fmt.Errorf("failed to parse message template: %w", errors.New("unknown channel 'Channel(0)'")),
		},
		{
			name: "ok",
			input: &entity.Message{
				Channel:         channel.Telegram,
				MessageTemplate: messagetemplate.Receipt,
				Params:          []byte(`{"orderId": 123, "commissionAmount": "1 KZT", "totalAmount": "1001 KZT"}`),
			},
			expectedError: nil,
			expectedContent: `<b>Чек</b>

Заказ <code>123</code> успешно оплачен

Комиссия: 1 KZT
Сумма к списанию: 1001 KZT

Спасибо за покупку`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := services.Message.getContentFromTemplate(tc.input)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedContent, data)
		})
	}
}

func TestMessage_appendQueueChannelMessage(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)

	cases := []struct {
		name          string
		channel       string
		messageID     string
		expected      string
		expectedError error
	}{
		{
			name:      "ok",
			channel:   fmt.Sprintf("ns::%d", channel.Telegram),
			messageID: "123",
			expected:  "123",
		},
		{
			name:          "error channel",
			channel:       "",
			messageID:     "123",
			expected:      "123",
			expectedError: InvalidChannelErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ch, _ := services.Message.getChannelNameFromPubSubKey(tc.channel)

			if tc.expectedError != nil {
				mocked.Logger.EXPECT().Error("get channel from PubSub key", "channel", tc.channel, "error", tc.expectedError)
			}

			go func() {
				for {
					assert.Equal(t, tc.expected, <-services.Message.chQueueChannels[ch])
				}
			}()

			services.Message.appendQueueChannelMessage(tc.channel, tc.messageID)
		})
	}
}

func mock(t *testing.T, mocked *utils.MockedInstances) *Store {
	t.Helper()

	repo := &repository.Store{Message: mocked.RepositoryMessage, User: mocked.RepositoryUser}
	channels := &channel.Store{Drivers: map[channel.Channel]channel.Driver{channel.Mock: mocked.ChannelDriver}}

	return NewStore(mocked.Logger, repo, channels)
}
