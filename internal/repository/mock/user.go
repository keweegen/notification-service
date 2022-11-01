// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	channel "github.com/keweegen/notification/internal/channel"
	entity "github.com/keweegen/notification/internal/entity"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateOrUpdateChannels mocks base method.
func (m *MockUser) CreateOrUpdateChannels(ctx context.Context, channels entity.UserChannels) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateChannels", ctx, channels)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateChannels indicates an expected call of CreateOrUpdateChannels.
func (mr *MockUserMockRecorder) CreateOrUpdateChannels(ctx, channels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateChannels", reflect.TypeOf((*MockUser)(nil).CreateOrUpdateChannels), ctx, channels)
}

// Exists mocks base method.
func (m *MockUser) Exists(ctx context.Context, userID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockUserMockRecorder) Exists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockUser)(nil).Exists), ctx, userID)
}

// FindByChannel mocks base method.
func (m *MockUser) FindByChannel(ctx context.Context, userID int64, channel channel.Channel) (*entity.UserChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByChannel", ctx, userID, channel)
	ret0, _ := ret[0].(*entity.UserChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByChannel indicates an expected call of FindByChannel.
func (mr *MockUserMockRecorder) FindByChannel(ctx, userID, channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByChannel", reflect.TypeOf((*MockUser)(nil).FindByChannel), ctx, userID, channel)
}
