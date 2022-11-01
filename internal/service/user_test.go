package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/keweegen/notification/internal/channel"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/repository"
	"github.com/keweegen/notification/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_CreateNotificationChannel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()

	cases := []struct {
		name          string
		userChannel   *entity.UserChannel
		expectedError error
	}{
		{
			name: "ok",
			userChannel: &entity.UserChannel{
				ID:        1,
				UserID:    1,
				Channel:   channel.Mock,
				Recipient: "mock",
				CanNotify: true,
			},
		},
		{
			name: "invalid channel",
			userChannel: &entity.UserChannel{
				ID:        1,
				UserID:    1,
				Channel:   channel.Channel(0),
				Recipient: "mock",
				CanNotify: true,
			},
			expectedError: InvalidChannelErr,
		},
		{
			name: "database error",
			userChannel: &entity.UserChannel{
				ID:        1,
				UserID:    1,
				Channel:   channel.Mock,
				Recipient: "mock",
				CanNotify: true,
			},
			expectedError: utils.FakeDatabaseError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !errors.Is(c.expectedError, InvalidChannelErr) {
				mocked.RepositoryUser.EXPECT().CreateChannel(ctx, c.userChannel).Return(c.expectedError)
			}

			err := services.User.CreateNotificationChannel(ctx, c.userChannel)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestUser_FindNotificationChannel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()
	userChannel := &entity.UserChannel{
		ID:        1,
		UserID:    1,
		Channel:   channel.Mock,
		Recipient: "mock",
		CanNotify: true,
	}

	cases := []struct {
		name                string
		userChannelID       int64
		expectedUserChannel *entity.UserChannel
		expectedError       error
	}{
		{
			name:                "ok",
			expectedUserChannel: userChannel,
		},
		{
			name:          "database error",
			userChannelID: userChannel.UserID,
			expectedError: utils.FakeDatabaseError,
		},
		{
			name:          "user channel not found error",
			userChannelID: userChannel.UserID,
			expectedError: repository.UserChannelNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mocked.RepositoryUser.EXPECT().FindChannel(ctx, c.userChannelID).Return(c.expectedUserChannel, c.expectedError)

			uc, err := services.User.FindNotificationChannel(ctx, c.userChannelID)
			assert.Equal(t, c.expectedUserChannel, uc)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestUser_UpdateChannel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()
	userChannelID := int64(1)

	cases := []struct {
		name                    string
		recipient               string
		canNotify               bool
		expectedFindUserChannel *entity.UserChannel
		expectedError           error
	}{
		{
			name:      "ok",
			recipient: "321",
			canNotify: true,
			expectedFindUserChannel: &entity.UserChannel{
				ID:        userChannelID,
				UserID:    1,
				Channel:   channel.Mock,
				Recipient: "123",
				CanNotify: false,
			},
		},
		{
			name:          "database error",
			recipient:     "321",
			canNotify:     true,
			expectedError: utils.FakeDatabaseError,
			expectedFindUserChannel: &entity.UserChannel{
				ID:        userChannelID,
				UserID:    1,
				Channel:   channel.Mock,
				Recipient: "123",
				CanNotify: false,
			},
		},
		{
			name:          "user channel not found error",
			expectedError: repository.UserChannelNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if errors.Is(c.expectedError, repository.UserChannelNotFound) {
				mocked.RepositoryUser.EXPECT().FindChannel(ctx, userChannelID).Return(c.expectedFindUserChannel, c.expectedError)
			} else {
				mocked.RepositoryUser.EXPECT().FindChannel(ctx, userChannelID).Return(c.expectedFindUserChannel, nil)

				c.expectedFindUserChannel.Recipient = c.recipient
				c.expectedFindUserChannel.CanNotify = c.canNotify
				mocked.RepositoryUser.EXPECT().UpdateChannel(ctx, c.expectedFindUserChannel).Return(c.expectedError)
			}

			err := services.User.UpdateNotificationChannel(ctx, userChannelID, c.recipient, c.canNotify)
			assert.Equal(t, c.expectedError, err)

			if c.expectedFindUserChannel != nil {
				assert.Equal(t, c.recipient, c.expectedFindUserChannel.Recipient)
				assert.Equal(t, c.canNotify, c.expectedFindUserChannel.CanNotify)
			}
		})
	}
}

func TestUser_DestroyChannel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()
	userChannelID := int64(1)

	cases := []struct {
		name          string
		expectedError error
	}{
		{
			name: "ok",
		},
		{
			name:          "database error",
			expectedError: utils.FakeDatabaseError,
		},
		{
			name:          "user channel not found error",
			expectedError: repository.UserChannelNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mocked.RepositoryUser.EXPECT().DestroyChannel(ctx, userChannelID).Return(c.expectedError)

			err := services.User.DestroyNotificationChannel(ctx, userChannelID)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestUser_FindNotificationChannels(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()
	userChannel := &entity.UserChannel{
		ID:        1,
		UserID:    1,
		Channel:   channel.Mock,
		Recipient: "mock",
		CanNotify: true,
	}

	cases := []struct {
		name                 string
		userID               int64
		expectedUserChannels entity.UserChannels
		expectedError        error
	}{
		{
			name:                 "ok",
			expectedUserChannels: entity.UserChannels{userChannel},
		},
		{
			name:          "database error",
			userID:        userChannel.UserID,
			expectedError: utils.FakeDatabaseError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mocked.RepositoryUser.EXPECT().FindChannelsByUser(ctx, c.userID).Return(c.expectedUserChannels, c.expectedError)

			uc, err := services.User.FindNotificationChannels(ctx, c.userID)
			assert.Equal(t, c.expectedUserChannels, uc)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
