// Code generated by MockGen. DO NOT EDIT.
// Source: node_label_module_version_reconciler.go

// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MocknodeLabelModuleVersionHelperAPI is a mock of nodeLabelModuleVersionHelperAPI interface.
type MocknodeLabelModuleVersionHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MocknodeLabelModuleVersionHelperAPIMockRecorder
}

// MocknodeLabelModuleVersionHelperAPIMockRecorder is the mock recorder for MocknodeLabelModuleVersionHelperAPI.
type MocknodeLabelModuleVersionHelperAPIMockRecorder struct {
	mock *MocknodeLabelModuleVersionHelperAPI
}

// NewMocknodeLabelModuleVersionHelperAPI creates a new mock instance.
func NewMocknodeLabelModuleVersionHelperAPI(ctrl *gomock.Controller) *MocknodeLabelModuleVersionHelperAPI {
	mock := &MocknodeLabelModuleVersionHelperAPI{ctrl: ctrl}
	mock.recorder = &MocknodeLabelModuleVersionHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknodeLabelModuleVersionHelperAPI) EXPECT() *MocknodeLabelModuleVersionHelperAPIMockRecorder {
	return m.recorder
}

// getLabelsPerModules mocks base method.
func (m *MocknodeLabelModuleVersionHelperAPI) getLabelsPerModules(ctx context.Context, nodeLabels map[string]string) map[string]*modulesVersionLabels {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLabelsPerModules", ctx, nodeLabels)
	ret0, _ := ret[0].(map[string]*modulesVersionLabels)
	return ret0
}

// getLabelsPerModules indicates an expected call of getLabelsPerModules.
func (mr *MocknodeLabelModuleVersionHelperAPIMockRecorder) getLabelsPerModules(ctx, nodeLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLabelsPerModules", reflect.TypeOf((*MocknodeLabelModuleVersionHelperAPI)(nil).getLabelsPerModules), ctx, nodeLabels)
}

// getModuleLoaderAndDevicePluginPods mocks base method.
func (m *MocknodeLabelModuleVersionHelperAPI) getModuleLoaderAndDevicePluginPods(ctx context.Context, nodeName string) ([]v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getModuleLoaderAndDevicePluginPods", ctx, nodeName)
	ret0, _ := ret[0].([]v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getModuleLoaderAndDevicePluginPods indicates an expected call of getModuleLoaderAndDevicePluginPods.
func (mr *MocknodeLabelModuleVersionHelperAPIMockRecorder) getModuleLoaderAndDevicePluginPods(ctx, nodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getModuleLoaderAndDevicePluginPods", reflect.TypeOf((*MocknodeLabelModuleVersionHelperAPI)(nil).getModuleLoaderAndDevicePluginPods), ctx, nodeName)
}

// reconcileLabels mocks base method.
func (m *MocknodeLabelModuleVersionHelperAPI) reconcileLabels(modulesLabels map[string]*modulesVersionLabels, moduleLoaderPods []v1.Pod) *reconcileLabelsResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "reconcileLabels", modulesLabels, moduleLoaderPods)
	ret0, _ := ret[0].(*reconcileLabelsResult)
	return ret0
}

// reconcileLabels indicates an expected call of reconcileLabels.
func (mr *MocknodeLabelModuleVersionHelperAPIMockRecorder) reconcileLabels(modulesLabels, moduleLoaderPods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "reconcileLabels", reflect.TypeOf((*MocknodeLabelModuleVersionHelperAPI)(nil).reconcileLabels), modulesLabels, moduleLoaderPods)
}

// updateNodeLabels mocks base method.
func (m *MocknodeLabelModuleVersionHelperAPI) updateNodeLabels(ctx context.Context, nodeName string, reconLabelsRes *reconcileLabelsResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "updateNodeLabels", ctx, nodeName, reconLabelsRes)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateNodeLabels indicates an expected call of updateNodeLabels.
func (mr *MocknodeLabelModuleVersionHelperAPIMockRecorder) updateNodeLabels(ctx, nodeName, reconLabelsRes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateNodeLabels", reflect.TypeOf((*MocknodeLabelModuleVersionHelperAPI)(nil).updateNodeLabels), ctx, nodeName, reconLabelsRes)
}
