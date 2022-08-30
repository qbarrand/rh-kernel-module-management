// Code generated by MockGen. DO NOT EDIT.
// Source: maker.go

// Package job is a generated GoMock package.
package job

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	v1 "k8s.io/api/batch/v1"
)

// MockMaker is a mock of Maker interface.
type MockMaker struct {
	ctrl     *gomock.Controller
	recorder *MockMakerMockRecorder
}

// MockMakerMockRecorder is the mock recorder for MockMaker.
type MockMakerMockRecorder struct {
	mock *MockMaker
}

// NewMockMaker creates a new mock instance.
func NewMockMaker(ctrl *gomock.Controller) *MockMaker {
	mock := &MockMaker{ctrl: ctrl}
	mock.recorder = &MockMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaker) EXPECT() *MockMakerMockRecorder {
	return m.recorder
}

// MakeJob mocks base method.
func (m *MockMaker) MakeJob(mod v1beta1.Module, buildConfig *v1beta1.Build, targetKernel, containerImage string) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeJob", mod, buildConfig, targetKernel, containerImage)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeJob indicates an expected call of MakeJob.
func (mr *MockMakerMockRecorder) MakeJob(mod, buildConfig, targetKernel, containerImage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJob", reflect.TypeOf((*MockMaker)(nil).MakeJob), mod, buildConfig, targetKernel, containerImage)
}
