// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_kubernetes_apps is a generated GoMock package.
package mock_kubernetes_apps

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockDeploymentClient is a mock of DeploymentClient interface.
type MockDeploymentClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentClientMockRecorder
}

// MockDeploymentClientMockRecorder is the mock recorder for MockDeploymentClient.
type MockDeploymentClientMockRecorder struct {
	mock *MockDeploymentClient
}

// NewMockDeploymentClient creates a new mock instance.
func NewMockDeploymentClient(ctrl *gomock.Controller) *MockDeploymentClient {
	mock := &MockDeploymentClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentClient) EXPECT() *MockDeploymentClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDeploymentClient) Get(ctx context.Context, objectKey client.ObjectKey) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objectKey)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDeploymentClientMockRecorder) Get(ctx, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentClient)(nil).Get), ctx, objectKey)
}

// List mocks base method.
func (m *MockDeploymentClient) List(ctx context.Context, options ...client.ListOption) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDeploymentClientMockRecorder) List(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeploymentClient)(nil).List), varargs...)
}

// MockReplicaSetClient is a mock of ReplicaSetClient interface.
type MockReplicaSetClient struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetClientMockRecorder
}

// MockReplicaSetClientMockRecorder is the mock recorder for MockReplicaSetClient.
type MockReplicaSetClientMockRecorder struct {
	mock *MockReplicaSetClient
}

// NewMockReplicaSetClient creates a new mock instance.
func NewMockReplicaSetClient(ctrl *gomock.Controller) *MockReplicaSetClient {
	mock := &MockReplicaSetClient{ctrl: ctrl}
	mock.recorder = &MockReplicaSetClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetClient) EXPECT() *MockReplicaSetClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReplicaSetClient) Get(ctx context.Context, objectKey client.ObjectKey) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objectKey)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReplicaSetClientMockRecorder) Get(ctx, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReplicaSetClient)(nil).Get), ctx, objectKey)
}
