// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Logger)

package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return _m.recorder
}

// Debug mocks base method
func (_m *MockLogger) Debug(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debug", _s...)
}

// Debug indicates an expected call of Debug
func (_mr *MockLoggerMockRecorder) Debug(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debug", arg0...)
}

// Debugf mocks base method
func (_m *MockLogger) Debugf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debugf", _s...)
}

// Debugf indicates an expected call of Debugf
func (_mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debugf", _s...)
}

// Error mocks base method
func (_m *MockLogger) Error(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Error", _s...)
}

// Error indicates an expected call of Error
func (_mr *MockLoggerMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Error", arg0...)
}

// Errorf mocks base method
func (_m *MockLogger) Errorf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Errorf", _s...)
}

// Errorf indicates an expected call of Errorf
func (_mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Errorf", _s...)
}

// Fatal mocks base method
func (_m *MockLogger) Fatal(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatal", _s...)
}

// Fatal indicates an expected call of Fatal
func (_mr *MockLoggerMockRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatal", arg0...)
}

// Fatalf mocks base method
func (_m *MockLogger) Fatalf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatalf", _s...)
}

// Fatalf indicates an expected call of Fatalf
func (_mr *MockLoggerMockRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatalf", _s...)
}

// Info mocks base method
func (_m *MockLogger) Info(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Info", _s...)
}

// Info indicates an expected call of Info
func (_mr *MockLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0...)
}

// Infof mocks base method
func (_m *MockLogger) Infof(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Infof", _s...)
}

// Infof indicates an expected call of Infof
func (_mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Infof", _s...)
}

// Warning mocks base method
func (_m *MockLogger) Warning(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warning", _s...)
}

// Warning indicates an expected call of Warning
func (_mr *MockLoggerMockRecorder) Warning(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warning", arg0...)
}

// Warningf mocks base method
func (_m *MockLogger) Warningf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warningf", _s...)
}

// Warningf indicates an expected call of Warningf
func (_mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warningf", _s...)
}
