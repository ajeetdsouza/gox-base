// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_gox is a generated GoMock package.
package mock_gox

import (
	http "net/http"
	reflect "reflect"
	time "time"

	metrics "github.com/devlibx/gox-base/metrics"
	gomock "github.com/golang/mock/gomock"
	zap "go.uber.org/zap"
)

// MockTimeService is a mock of TimeService interface.
type MockTimeService struct {
	ctrl     *gomock.Controller
	recorder *MockTimeServiceMockRecorder
}

// MockTimeServiceMockRecorder is the mock recorder for MockTimeService.
type MockTimeServiceMockRecorder struct {
	mock *MockTimeService
}

// NewMockTimeService creates a new mock instance.
func NewMockTimeService(ctrl *gomock.Controller) *MockTimeService {
	mock := &MockTimeService{ctrl: ctrl}
	mock.recorder = &MockTimeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeService) EXPECT() *MockTimeServiceMockRecorder {
	return m.recorder
}

// Now mocks base method.
func (m *MockTimeService) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockTimeServiceMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockTimeService)(nil).Now))
}

// Sleep mocks base method.
func (m *MockTimeService) Sleep(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep.
func (mr *MockTimeServiceMockRecorder) Sleep(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockTimeService)(nil).Sleep), d)
}

// MockCrossFunction is a mock of CrossFunction interface.
type MockCrossFunction struct {
	ctrl     *gomock.Controller
	recorder *MockCrossFunctionMockRecorder
}

// MockCrossFunctionMockRecorder is the mock recorder for MockCrossFunction.
type MockCrossFunctionMockRecorder struct {
	mock *MockCrossFunction
}

// NewMockCrossFunction creates a new mock instance.
func NewMockCrossFunction(ctrl *gomock.Controller) *MockCrossFunction {
	mock := &MockCrossFunction{ctrl: ctrl}
	mock.recorder = &MockCrossFunctionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossFunction) EXPECT() *MockCrossFunctionMockRecorder {
	return m.recorder
}

// Counter mocks base method.
func (m *MockCrossFunction) Counter(counterName string) metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Counter", counterName)
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// Counter indicates an expected call of Counter.
func (mr *MockCrossFunctionMockRecorder) Counter(counterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counter", reflect.TypeOf((*MockCrossFunction)(nil).Counter), counterName)
}

// HttpHandler mocks base method.
func (m *MockCrossFunction) HttpHandler() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpHandler")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// HttpHandler indicates an expected call of HttpHandler.
func (mr *MockCrossFunctionMockRecorder) HttpHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpHandler", reflect.TypeOf((*MockCrossFunction)(nil).HttpHandler))
}

// Initialize mocks base method.
func (m *MockCrossFunction) Initialize(configuration metrics.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", configuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockCrossFunctionMockRecorder) Initialize(configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockCrossFunction)(nil).Initialize), configuration)
}

// Logger mocks base method.
func (m *MockCrossFunction) Logger() *zap.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(*zap.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockCrossFunctionMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockCrossFunction)(nil).Logger))
}

// Now mocks base method.
func (m *MockCrossFunction) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockCrossFunctionMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockCrossFunction)(nil).Now))
}

// RegisterCounter mocks base method.
func (m *MockCrossFunction) RegisterCounter(name, help string, labels []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCounter", name, help, labels)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCounter indicates an expected call of RegisterCounter.
func (mr *MockCrossFunctionMockRecorder) RegisterCounter(name, help, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCounter", reflect.TypeOf((*MockCrossFunction)(nil).RegisterCounter), name, help, labels)
}

// Sleep mocks base method.
func (m *MockCrossFunction) Sleep(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep.
func (mr *MockCrossFunctionMockRecorder) Sleep(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockCrossFunction)(nil).Sleep), d)
}
