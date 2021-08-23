// Code generated by MockGen. DO NOT EDIT.
// Source: magma/lte/gateway/log (interfaces: Logger)

// Package log is a generated GoMock package.
package log

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug() Printer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(Printer)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug))
}

// Error mocks base method.
func (m *MockLogger) Error() Printer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(Printer)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error))
}

// Info mocks base method.
func (m *MockLogger) Info() Printer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(Printer)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info))
}

// Level mocks base method.
func (m *MockLogger) Level() Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(Level)
	return ret0
}

// Level indicates an expected call of Level.
func (mr *MockLoggerMockRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockLogger)(nil).Level))
}

// Named mocks base method.
func (m *MockLogger) Named(arg0 string) Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Named", arg0)
	ret0, _ := ret[0].(Logger)
	return ret0
}

// Named indicates an expected call of Named.
func (mr *MockLoggerMockRecorder) Named(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Named", reflect.TypeOf((*MockLogger)(nil).Named), arg0)
}

// SetLevel mocks base method.
func (m *MockLogger) SetLevel(arg0 Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLevel", arg0)
}

// SetLevel indicates an expected call of SetLevel.
func (mr *MockLoggerMockRecorder) SetLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLevel", reflect.TypeOf((*MockLogger)(nil).SetLevel), arg0)
}

// Warning mocks base method.
func (m *MockLogger) Warning() Printer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Warning")
	ret0, _ := ret[0].(Printer)
	return ret0
}

// Warning indicates an expected call of Warning.
func (mr *MockLoggerMockRecorder) Warning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockLogger)(nil).Warning))
}

// With mocks base method.
func (m *MockLogger) With(arg0 string, arg1 interface{}) Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", arg0, arg1)
	ret0, _ := ret[0].(Logger)
	return ret0
}

// With indicates an expected call of With.
func (mr *MockLoggerMockRecorder) With(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockLogger)(nil).With), arg0, arg1)
}
