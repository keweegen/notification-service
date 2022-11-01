package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/utils"
	"testing"
)

func TestMessageChecker_Do_OK(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()

	message := mocked.FakeMessage()
	messages := entity.Messages{message}
	userChannel := mocked.FakeUserChannel()

	mocked.Logger.EXPECT().Debug("get process messages")
	mocked.Logger.EXPECT().Debug("sending message")
	mocked.Logger.EXPECT().Debug("message sent", "channel", message.Channel, "content", gomock.Any())

	mocked.RepositoryMessage.EXPECT().FindProcessMessages(ctx, gomock.Any(), gomock.Any()).Return(messages, nil)
	mocked.RepositoryMessage.EXPECT().CreateStatus(ctx, message.ID, entity.MessageStatusSending, "Sending a message").Return(nil)
	mocked.RepositoryUser.EXPECT().FindByChannel(ctx, message.UserID, message.Channel).Return(userChannel, nil)
	mocked.ChannelDriver.EXPECT().Send(userChannel.Recipient, gomock.Any()).Return(nil)
	mocked.RepositoryMessage.EXPECT().CreateStatus(ctx, message.ID, entity.MessageStatusSent, "Message sent").Return(nil)

	go services.MessageChecker.Do(ctx, mocked.QuitCh)

	mocked.WriteQuitChannel()
}

func TestMessageChecker_Do_ErrorGetProcessMessages(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()
	expectedError := errors.New("database error :(")

	mocked.Logger.EXPECT().Debug("get process messages")
	mocked.RepositoryMessage.EXPECT().FindProcessMessages(ctx, gomock.Any(), gomock.Any()).Return(nil, expectedError)
	mocked.Logger.EXPECT().Error("failed to get process messages",
		"error", fmt.Errorf("get process messages: %w", expectedError))

	go services.MessageChecker.Do(ctx, mocked.QuitCh)

	mocked.WriteQuitChannel()
}

func TestMessageChecker_Do_ErrorResendMessages(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mocked := utils.NewMockedInstances(controller)
	mocked.ExpectLoggerWithServices()

	services := mock(t, mocked)
	ctx := context.Background()

	message := mocked.FakeMessage()
	messages := entity.Messages{message}
	expectedError := errors.New("database error :(")

	mocked.Logger.EXPECT().Debug("get process messages")
	mocked.Logger.EXPECT().Debug("sending message")
	mocked.Logger.EXPECT().Error("failed to send message",
		"messageId", message.ID,
		"error", fmt.Errorf("find user notification channel: %w", expectedError))

	mocked.RepositoryMessage.EXPECT().FindProcessMessages(ctx, gomock.Any(), gomock.Any()).Return(messages, nil)
	mocked.RepositoryMessage.EXPECT().CreateStatus(ctx, message.ID, entity.MessageStatusSending, "Sending a message").Return(nil)
	mocked.RepositoryUser.EXPECT().FindByChannel(ctx, message.UserID, message.Channel).Return(nil, expectedError)

	go services.MessageChecker.Do(ctx, mocked.QuitCh)

	mocked.WriteQuitChannel()
}
