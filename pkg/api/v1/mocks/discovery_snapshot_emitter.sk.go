// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/discovery_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	v10 "github.com/solo-io/supergloo/pkg/api/v1"
)

// MockDiscoveryEmitter is a mock of DiscoveryEmitter interface
type MockDiscoveryEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryEmitterMockRecorder
}

// MockDiscoveryEmitterMockRecorder is the mock recorder for MockDiscoveryEmitter
type MockDiscoveryEmitterMockRecorder struct {
	mock *MockDiscoveryEmitter
}

// NewMockDiscoveryEmitter creates a new mock instance
func NewMockDiscoveryEmitter(ctrl *gomock.Controller) *MockDiscoveryEmitter {
	mock := &MockDiscoveryEmitter{ctrl: ctrl}
	mock.recorder = &MockDiscoveryEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoveryEmitter) EXPECT() *MockDiscoveryEmitterMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockDiscoveryEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockDiscoveryEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDiscoveryEmitter)(nil).Register))
}

// Pod mocks base method
func (m *MockDiscoveryEmitter) Pod() kubernetes.PodClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod")
	ret0, _ := ret[0].(kubernetes.PodClient)
	return ret0
}

// Pod indicates an expected call of Pod
func (mr *MockDiscoveryEmitterMockRecorder) Pod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockDiscoveryEmitter)(nil).Pod))
}

// Upstream mocks base method
func (m *MockDiscoveryEmitter) Upstream() v1.UpstreamClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upstream")
	ret0, _ := ret[0].(v1.UpstreamClient)
	return ret0
}

// Upstream indicates an expected call of Upstream
func (mr *MockDiscoveryEmitterMockRecorder) Upstream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upstream", reflect.TypeOf((*MockDiscoveryEmitter)(nil).Upstream))
}

// TlsSecret mocks base method
func (m *MockDiscoveryEmitter) TlsSecret() v10.TlsSecretClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TlsSecret")
	ret0, _ := ret[0].(v10.TlsSecretClient)
	return ret0
}

// TlsSecret indicates an expected call of TlsSecret
func (mr *MockDiscoveryEmitterMockRecorder) TlsSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsSecret", reflect.TypeOf((*MockDiscoveryEmitter)(nil).TlsSecret))
}

// Deployment mocks base method
func (m *MockDiscoveryEmitter) Deployment() kubernetes.DeploymentClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployment")
	ret0, _ := ret[0].(kubernetes.DeploymentClient)
	return ret0
}

// Deployment indicates an expected call of Deployment
func (mr *MockDiscoveryEmitterMockRecorder) Deployment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockDiscoveryEmitter)(nil).Deployment))
}

// Snapshots mocks base method
func (m *MockDiscoveryEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v10.DiscoverySnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v10.DiscoverySnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockDiscoveryEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockDiscoveryEmitter)(nil).Snapshots), watchNamespaces, opts)
}
