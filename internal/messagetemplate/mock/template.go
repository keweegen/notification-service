// Code generated by MockGen. DO NOT EDIT.
// Source: template.go

// Package mock_messagetemplate is a generated GoMock package.
package mock_messagetemplate

import (
	template "html/template"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/volatiletech/sqlboiler/v4/types"
)

// MockTemplate is a mock of Template interface.
type MockTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockTemplateMockRecorder
}

// MockTemplateMockRecorder is the mock recorder for MockTemplate.
type MockTemplateMockRecorder struct {
	mock *MockTemplate
}

// NewMockTemplate creates a new mock instance.
func NewMockTemplate(ctrl *gomock.Controller) *MockTemplate {
	mock := &MockTemplate{ctrl: ctrl}
	mock.recorder = &MockTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemplate) EXPECT() *MockTemplateMockRecorder {
	return m.recorder
}

// EmailTemplate mocks base method.
func (m *MockTemplate) EmailTemplate() *template.Template {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmailTemplate")
	ret0, _ := ret[0].(*template.Template)
	return ret0
}

// EmailTemplate indicates an expected call of EmailTemplate.
func (mr *MockTemplateMockRecorder) EmailTemplate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmailTemplate", reflect.TypeOf((*MockTemplate)(nil).EmailTemplate))
}

// Name mocks base method.
func (m *MockTemplate) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockTemplateMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockTemplate)(nil).Name))
}

// SetParams mocks base method.
func (m *MockTemplate) SetParams(data types.JSON) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetParams", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParams indicates an expected call of SetParams.
func (mr *MockTemplateMockRecorder) SetParams(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParams", reflect.TypeOf((*MockTemplate)(nil).SetParams), data)
}

// TelegramTemplate mocks base method.
func (m *MockTemplate) TelegramTemplate() *template.Template {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TelegramTemplate")
	ret0, _ := ret[0].(*template.Template)
	return ret0
}

// TelegramTemplate indicates an expected call of TelegramTemplate.
func (mr *MockTemplateMockRecorder) TelegramTemplate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TelegramTemplate", reflect.TypeOf((*MockTemplate)(nil).TelegramTemplate))
}
