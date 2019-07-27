// Code generated by MockGen. DO NOT EDIT.
// Source: BudgetService.go

// Package mock is a generated GoMock package.
package logic

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIBudgetRepository is a mock of IBudgetRepository interface
type MockIBudgetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBudgetRepositoryMockRecorder
}

// MockIBudgetRepositoryMockRecorder is the mock recorder for MockIBudgetRepository
type MockIBudgetRepositoryMockRecorder struct {
	mock *MockIBudgetRepository
}

// NewMockIBudgetRepository creates a new mock instance
func NewMockIBudgetRepository(ctrl *gomock.Controller) *MockIBudgetRepository {
	mock := &MockIBudgetRepository{ctrl: ctrl}
	mock.recorder = &MockIBudgetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIBudgetRepository) EXPECT() *MockIBudgetRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockIBudgetRepository) GetAll() []Bugget {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]Bugget)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockIBudgetRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIBudgetRepository)(nil).GetAll))
}