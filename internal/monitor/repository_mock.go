// Code generated by MockGen. DO NOT EDIT.
// Source: internal/monitor/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/monitor/repository.go -destination=internal/monitor/mocks/repository_mock.go -package=monitor
//

// Package monitor is a generated GoMock package.
package monitor

import (
	context "context"
	reflect "reflect"

	probe "github.com/RykoL/uptime-probe/internal/monitor/probe"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetMonitors mocks base method.
func (m *MockRepository) GetMonitors(ctx context.Context) ([]*Monitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonitors", ctx)
	ret0, _ := ret[0].([]*Monitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonitors indicates an expected call of GetMonitors.
func (mr *MockRepositoryMockRecorder) GetMonitors(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonitors", reflect.TypeOf((*MockRepository)(nil).GetMonitors), ctx)
}

// RecordProbeResult mocks base method.
func (m *MockRepository) RecordProbeResult(ctx context.Context, monitorId int, result *probe.ProbeResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordProbeResult", ctx, monitorId, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordProbeResult indicates an expected call of RecordProbeResult.
func (mr *MockRepositoryMockRecorder) RecordProbeResult(ctx, monitorId, result any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordProbeResult", reflect.TypeOf((*MockRepository)(nil).RecordProbeResult), ctx, monitorId, result)
}

// SaveMonitor mocks base method.
func (m *MockRepository) SaveMonitor(ctx context.Context, arg1 *Monitor) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMonitor", ctx, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveMonitor indicates an expected call of SaveMonitor.
func (mr *MockRepositoryMockRecorder) SaveMonitor(ctx, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMonitor", reflect.TypeOf((*MockRepository)(nil).SaveMonitor), ctx, arg1)
}
