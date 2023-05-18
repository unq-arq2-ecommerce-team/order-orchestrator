// Code generated by MockGen. DO NOT EDIT.
// Source: notification.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/unq-arq2-ecommerce-team/order-orchestrator/src/domain/model"
)

// MockNotificationRepository is a mock of NotificationRepository interface.
type MockNotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRepositoryMockRecorder
}

// MockNotificationRepositoryMockRecorder is the mock recorder for MockNotificationRepository.
type MockNotificationRepositoryMockRecorder struct {
	mock *MockNotificationRepository
}

// NewMockNotificationRepository creates a new mock instance.
func NewMockNotificationRepository(ctrl *gomock.Controller) *MockNotificationRepository {
	mock := &MockNotificationRepository{ctrl: ctrl}
	mock.recorder = &MockNotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRepository) EXPECT() *MockNotificationRepositoryMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockNotificationRepository) Send(ctx context.Context, notification model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNotificationRepositoryMockRecorder) Send(ctx, notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotificationRepository)(nil).Send), ctx, notification)
}
