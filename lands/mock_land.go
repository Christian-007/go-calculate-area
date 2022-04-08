// Code generated by MockGen. DO NOT EDIT.
// Source: land.go

// Package mock_lands is a generated GoMock package.
package lands

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLand is a mock of Land interface.
type MockLand struct {
	ctrl     *gomock.Controller
	recorder *MockLandMockRecorder
}

// MockLandMockRecorder is the mock recorder for MockLand.
type MockLandMockRecorder struct {
	mock *MockLand
}

// NewMockLand creates a new mock instance.
func NewMockLand(ctrl *gomock.Controller) *MockLand {
	mock := &MockLand{ctrl: ctrl}
	mock.recorder = &MockLandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLand) EXPECT() *MockLandMockRecorder {
	return m.recorder
}

// CalculateArea mocks base method.
func (m *MockLand) CalculateArea() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateArea")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateArea indicates an expected call of CalculateArea.
func (mr *MockLandMockRecorder) CalculateArea() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateArea", reflect.TypeOf((*MockLand)(nil).CalculateArea))
}
