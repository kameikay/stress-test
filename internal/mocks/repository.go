// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/ports.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStressTestRepository is a mock of StressTestRepository interface.
type MockStressTestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStressTestRepositoryMockRecorder
}

// MockStressTestRepositoryMockRecorder is the mock recorder for MockStressTestRepository.
type MockStressTestRepositoryMockRecorder struct {
	mock *MockStressTestRepository
}

// NewMockStressTestRepository creates a new mock instance.
func NewMockStressTestRepository(ctrl *gomock.Controller) *MockStressTestRepository {
	mock := &MockStressTestRepository{ctrl: ctrl}
	mock.recorder = &MockStressTestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStressTestRepository) EXPECT() *MockStressTestRepositoryMockRecorder {
	return m.recorder
}

// GetErrorRequests mocks base method.
func (m *MockStressTestRepository) GetErrorRequests() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorRequests")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetErrorRequests indicates an expected call of GetErrorRequests.
func (mr *MockStressTestRepositoryMockRecorder) GetErrorRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorRequests", reflect.TypeOf((*MockStressTestRepository)(nil).GetErrorRequests))
}

// GetOthersStatus mocks base method.
func (m *MockStressTestRepository) GetOthersStatus() map[int]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOthersStatus")
	ret0, _ := ret[0].(map[int]int)
	return ret0
}

// GetOthersStatus indicates an expected call of GetOthersStatus.
func (mr *MockStressTestRepositoryMockRecorder) GetOthersStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOthersStatus", reflect.TypeOf((*MockStressTestRepository)(nil).GetOthersStatus))
}

// GetStatus200 mocks base method.
func (m *MockStressTestRepository) GetStatus200() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus200")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetStatus200 indicates an expected call of GetStatus200.
func (mr *MockStressTestRepositoryMockRecorder) GetStatus200() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus200", reflect.TypeOf((*MockStressTestRepository)(nil).GetStatus200))
}

// GetTotalRequests mocks base method.
func (m *MockStressTestRepository) GetTotalRequests() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequests")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetTotalRequests indicates an expected call of GetTotalRequests.
func (mr *MockStressTestRepositoryMockRecorder) GetTotalRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequests", reflect.TypeOf((*MockStressTestRepository)(nil).GetTotalRequests))
}

// Save mocks base method.
func (m *MockStressTestRepository) Save(statusCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Save", statusCode)
}

// Save indicates an expected call of Save.
func (mr *MockStressTestRepositoryMockRecorder) Save(statusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStressTestRepository)(nil).Save), statusCode)
}
