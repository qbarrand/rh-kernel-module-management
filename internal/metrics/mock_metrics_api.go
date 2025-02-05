// Code generated by MockGen. DO NOT EDIT.
// Source: metrics.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockMetrics) Register() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register")
}

// Register indicates an expected call of Register.
func (mr *MockMetricsMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockMetrics)(nil).Register))
}

// SetKMMDevicePluginNum mocks base method.
func (m *MockMetrics) SetKMMDevicePluginNum(value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMDevicePluginNum", value)
}

// SetKMMDevicePluginNum indicates an expected call of SetKMMDevicePluginNum.
func (mr *MockMetricsMockRecorder) SetKMMDevicePluginNum(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMDevicePluginNum", reflect.TypeOf((*MockMetrics)(nil).SetKMMDevicePluginNum), value)
}

// SetKMMInClusterBuildNum mocks base method.
func (m *MockMetrics) SetKMMInClusterBuildNum(value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMInClusterBuildNum", value)
}

// SetKMMInClusterBuildNum indicates an expected call of SetKMMInClusterBuildNum.
func (mr *MockMetricsMockRecorder) SetKMMInClusterBuildNum(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMInClusterBuildNum", reflect.TypeOf((*MockMetrics)(nil).SetKMMInClusterBuildNum), value)
}

// SetKMMInClusterSignNum mocks base method.
func (m *MockMetrics) SetKMMInClusterSignNum(value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMInClusterSignNum", value)
}

// SetKMMInClusterSignNum indicates an expected call of SetKMMInClusterSignNum.
func (mr *MockMetricsMockRecorder) SetKMMInClusterSignNum(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMInClusterSignNum", reflect.TypeOf((*MockMetrics)(nil).SetKMMInClusterSignNum), value)
}

// SetKMMModprobeArgs mocks base method.
func (m *MockMetrics) SetKMMModprobeArgs(modName, namespace, modprobeArgs string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMModprobeArgs", modName, namespace, modprobeArgs)
}

// SetKMMModprobeArgs indicates an expected call of SetKMMModprobeArgs.
func (mr *MockMetricsMockRecorder) SetKMMModprobeArgs(modName, namespace, modprobeArgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMModprobeArgs", reflect.TypeOf((*MockMetrics)(nil).SetKMMModprobeArgs), modName, namespace, modprobeArgs)
}

// SetKMMModprobeRawArgs mocks base method.
func (m *MockMetrics) SetKMMModprobeRawArgs(modName, namespace, modprobeArgs string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMModprobeRawArgs", modName, namespace, modprobeArgs)
}

// SetKMMModprobeRawArgs indicates an expected call of SetKMMModprobeRawArgs.
func (mr *MockMetricsMockRecorder) SetKMMModprobeRawArgs(modName, namespace, modprobeArgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMModprobeRawArgs", reflect.TypeOf((*MockMetrics)(nil).SetKMMModprobeRawArgs), modName, namespace, modprobeArgs)
}

// SetKMMModulesNum mocks base method.
func (m *MockMetrics) SetKMMModulesNum(value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMModulesNum", value)
}

// SetKMMModulesNum indicates an expected call of SetKMMModulesNum.
func (mr *MockMetricsMockRecorder) SetKMMModulesNum(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMModulesNum", reflect.TypeOf((*MockMetrics)(nil).SetKMMModulesNum), value)
}

// SetKMMPreflightsNum mocks base method.
func (m *MockMetrics) SetKMMPreflightsNum(value int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKMMPreflightsNum", value)
}

// SetKMMPreflightsNum indicates an expected call of SetKMMPreflightsNum.
func (mr *MockMetricsMockRecorder) SetKMMPreflightsNum(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKMMPreflightsNum", reflect.TypeOf((*MockMetrics)(nil).SetKMMPreflightsNum), value)
}
