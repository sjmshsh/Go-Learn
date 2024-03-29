// Code generated by MockGen. DO NOT EDIT.
// Source: ./mockdemo.go

// Package mockdemo is a generated GoMock package.
package mockdemo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMail is a mock of Mail interface.
type MockMail struct {
	ctrl     *gomock.Controller
	recorder *MockMailMockRecorder
}

// MockMailMockRecorder is the mock recorder for MockMail.
type MockMailMockRecorder struct {
	mock *MockMail
}

// NewMockMail creates a new mock instance.
func NewMockMail(ctrl *gomock.Controller) *MockMail {
	mock := &MockMail{ctrl: ctrl}
	mock.recorder = &MockMailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMail) EXPECT() *MockMailMockRecorder {
	return m.recorder
}

// sendMail mocks base method.
func (m *MockMail) sendMail(subject, sender, dst, body string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "sendMail", subject, sender, dst, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// sendMail indicates an expected call of sendMail.
func (mr *MockMailMockRecorder) sendMail(subject, sender, dst, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "sendMail", reflect.TypeOf((*MockMail)(nil).sendMail), subject, sender, dst, body)
}
