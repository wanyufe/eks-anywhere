// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/providers/cloudstack/internal/templates (interfaces: CloudMonkeyClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudMonkeyClient is a mock of CloudMonkeyClient interface.
type MockCloudMonkeyClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMonkeyClientMockRecorder
}

// MockCloudMonkeyClientMockRecorder is the mock recorder for MockCloudMonkeyClient.
type MockCloudMonkeyClientMockRecorder struct {
	mock *MockCloudMonkeyClient
}

// NewMockCloudMonkeyClient creates a new mock instance.
func NewMockCloudMonkeyClient(ctrl *gomock.Controller) *MockCloudMonkeyClient {
	mock := &MockCloudMonkeyClient{ctrl: ctrl}
	mock.recorder = &MockCloudMonkeyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudMonkeyClient) EXPECT() *MockCloudMonkeyClientMockRecorder {
	return m.recorder
}

// SearchComputeOffering mocks base method.
func (m *MockCloudMonkeyClient) SearchComputeOffering(arg0 context.Context, arg1, arg2, arg3, arg4 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchComputeOffering", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchComputeOffering indicates an expected call of SearchComputeOffering.
func (mr *MockCloudMonkeyClientMockRecorder) SearchComputeOffering(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchComputeOffering", reflect.TypeOf((*MockCloudMonkeyClient)(nil).SearchComputeOffering), arg0, arg1, arg2, arg3, arg4)
}

// SearchDiskOffering mocks base method.
func (m *MockCloudMonkeyClient) SearchDiskOffering(arg0 context.Context, arg1, arg2, arg3, arg4 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDiskOffering", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDiskOffering indicates an expected call of SearchDiskOffering.
func (mr *MockCloudMonkeyClientMockRecorder) SearchDiskOffering(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDiskOffering", reflect.TypeOf((*MockCloudMonkeyClient)(nil).SearchDiskOffering), arg0, arg1, arg2, arg3, arg4)
}

// SearchTemplate mocks base method.
func (m *MockCloudMonkeyClient) SearchTemplate(arg0 context.Context, arg1, arg2, arg3, arg4 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTemplate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTemplate indicates an expected call of SearchTemplate.
func (mr *MockCloudMonkeyClientMockRecorder) SearchTemplate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTemplate", reflect.TypeOf((*MockCloudMonkeyClient)(nil).SearchTemplate), arg0, arg1, arg2, arg3, arg4)
}

// ValidateCloudStackSetup mocks base method.
func (m *MockCloudMonkeyClient) ValidateCloudStackSetup(arg0 context.Context, arg1 *v1alpha1.CloudStackDeploymentConfig, arg2 *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCloudStackSetup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCloudStackSetup indicates an expected call of ValidateCloudStackSetup.
func (mr *MockCloudMonkeyClientMockRecorder) ValidateCloudStackSetup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCloudStackSetup", reflect.TypeOf((*MockCloudMonkeyClient)(nil).ValidateCloudStackSetup), arg0, arg1, arg2)
}

// ValidateCloudStackSetupMachineConfig mocks base method.
func (m *MockCloudMonkeyClient) ValidateCloudStackSetupMachineConfig(arg0 context.Context, arg1 *v1alpha1.CloudStackDeploymentConfig, arg2 *v1alpha1.CloudStackMachineConfig, arg3 *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCloudStackSetupMachineConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCloudStackSetupMachineConfig indicates an expected call of ValidateCloudStackSetupMachineConfig.
func (mr *MockCloudMonkeyClientMockRecorder) ValidateCloudStackSetupMachineConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCloudStackSetupMachineConfig", reflect.TypeOf((*MockCloudMonkeyClient)(nil).ValidateCloudStackSetupMachineConfig), arg0, arg1, arg2, arg3)
}
