// Code generated by MockGen. DO NOT EDIT.
// Source: local_job.go

// Package job is a generated GoMock package.
package job

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	events "github.com/openshift/assisted-service/internal/events"
	reflect "reflect"
)

// MockLocalJob is a mock of LocalJob interface
type MockLocalJob struct {
	ctrl     *gomock.Controller
	recorder *MockLocalJobMockRecorder
}

// MockLocalJobMockRecorder is the mock recorder for MockLocalJob
type MockLocalJobMockRecorder struct {
	mock *MockLocalJob
}

// NewMockLocalJob creates a new mock instance
func NewMockLocalJob(ctrl *gomock.Controller) *MockLocalJob {
	mock := &MockLocalJob{ctrl: ctrl}
	mock.recorder = &MockLocalJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalJob) EXPECT() *MockLocalJobMockRecorder {
	return m.recorder
}

// GenerateISO mocks base method
func (m *MockLocalJob) GenerateISO(ctx context.Context, cluster common.Cluster, jobName, imageName, ignitionConfig string, eventsHandler events.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateISO", ctx, cluster, jobName, imageName, ignitionConfig, eventsHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateISO indicates an expected call of GenerateISO
func (mr *MockLocalJobMockRecorder) GenerateISO(ctx, cluster, jobName, imageName, ignitionConfig, eventsHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateISO", reflect.TypeOf((*MockLocalJob)(nil).GenerateISO), ctx, cluster, jobName, imageName, ignitionConfig, eventsHandler)
}

// GenerateInstallConfig mocks base method
func (m *MockLocalJob) GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInstallConfig", ctx, cluster, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInstallConfig indicates an expected call of GenerateInstallConfig
func (mr *MockLocalJobMockRecorder) GenerateInstallConfig(ctx, cluster, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).GenerateInstallConfig), ctx, cluster, cfg)
}

// AbortInstallConfig mocks base method
func (m *MockLocalJob) AbortInstallConfig(ctx context.Context, cluster common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortInstallConfig", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortInstallConfig indicates an expected call of AbortInstallConfig
func (mr *MockLocalJobMockRecorder) AbortInstallConfig(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).AbortInstallConfig), ctx, cluster)
}
