// Code generated by MockGen. DO NOT EDIT.
// Source: helper.go

// Package nmc is a generated GoMock package.
package nmc

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	api "github.com/kubernetes-sigs/kernel-module-management/internal/api"
	gomock "go.uber.org/mock/gomock"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHelper) Get(ctx context.Context, name string) (*v1beta1.NodeModulesConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name)
	ret0, _ := ret[0].(*v1beta1.NodeModulesConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHelperMockRecorder) Get(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHelper)(nil).Get), ctx, name)
}

// GetModuleEntry mocks base method.
func (m *MockHelper) GetModuleEntry(nmc *v1beta1.NodeModulesConfig, modNamespace, modName string) (*v1beta1.NodeModuleSpec, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleEntry", nmc, modNamespace, modName)
	ret0, _ := ret[0].(*v1beta1.NodeModuleSpec)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetModuleEntry indicates an expected call of GetModuleEntry.
func (mr *MockHelperMockRecorder) GetModuleEntry(nmc, modNamespace, modName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleEntry", reflect.TypeOf((*MockHelper)(nil).GetModuleEntry), nmc, modNamespace, modName)
}

// RemoveModuleConfig mocks base method.
func (m *MockHelper) RemoveModuleConfig(nmc *v1beta1.NodeModulesConfig, namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveModuleConfig", nmc, namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveModuleConfig indicates an expected call of RemoveModuleConfig.
func (mr *MockHelperMockRecorder) RemoveModuleConfig(nmc, namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModuleConfig", reflect.TypeOf((*MockHelper)(nil).RemoveModuleConfig), nmc, namespace, name)
}

// SetModuleConfig mocks base method.
func (m *MockHelper) SetModuleConfig(nmc *v1beta1.NodeModulesConfig, mld *api.ModuleLoaderData, moduleConfig *v1beta1.ModuleConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModuleConfig", nmc, mld, moduleConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModuleConfig indicates an expected call of SetModuleConfig.
func (mr *MockHelperMockRecorder) SetModuleConfig(nmc, mld, moduleConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleConfig", reflect.TypeOf((*MockHelper)(nil).SetModuleConfig), nmc, mld, moduleConfig)
}
