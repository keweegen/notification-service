package utils

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/keweegen/notification/internal/channel"
	mockChannel "github.com/keweegen/notification/internal/channel/mock"
	"github.com/keweegen/notification/internal/entity"
	"github.com/keweegen/notification/internal/messagetemplate"
	mockRepository "github.com/keweegen/notification/internal/repository/mock"
	mockLogger "github.com/keweegen/notification/logger/mock"
	"time"
)

var FakeDatabaseError = errors.New("database error :(")

type MockedInstances struct {
	Context           context.Context
	Logger            *mockLogger.MockLogger
	ChannelDriver     *mockChannel.MockDriver
	RepositoryMessage *mockRepository.MockMessage
	RepositoryUser    *mockRepository.MockUser

	ctxCancel context.CancelFunc
}

func NewMockedInstances(controller *gomock.Controller) *MockedInstances {
	ctx, cancel := context.WithCancel(context.Background())

	return &MockedInstances{
		Logger:            mockLogger.NewMockLogger(controller),
		ChannelDriver:     mockChannel.NewMockDriver(controller),
		RepositoryMessage: mockRepository.NewMockMessage(controller),
		RepositoryUser:    mockRepository.NewMockUser(controller),
		Context:           ctx,
		ctxCancel:         cancel,
	}
}

func (m *MockedInstances) ExpectLoggerWithServices() {
	m.Logger.EXPECT().With("service", "message").Return(m.Logger)
	m.Logger.EXPECT().With("service", "messageChecker").Return(m.Logger)
}

func (m *MockedInstances) ContextCancel() {
	m.ctxCancel()
	time.Sleep(10 * time.Millisecond)
}

func (m *MockedInstances) FakeUserChannel() *entity.UserChannel {
	return &entity.UserChannel{
		UserID:    1234567890,
		Channel:   channel.Mock,
		Recipient: "408354752",
		CanNotify: true,
	}
}

func (m *MockedInstances) FakeMessage() *entity.Message {
	return &entity.Message{
		ID:              "NS-001-001-0000000001234567890-1666115824000-0000000001234567890",
		UserID:          1234567890,
		Channel:         channel.Mock,
		MessageTemplate: messagetemplate.Receipt,
		Timestamp:       time.Now().UnixMilli(),
		ExternalID:      1234567890,
		Params:          []byte(`{"orderId": 123, "commissionAmount": "1 KZT", "totalAmount": "1001 KZT"}`),
	}
}
