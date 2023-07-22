// Code generated by MockGen. DO NOT EDIT.
// Source: container.go

// Package builder is a generated GoMock package.
package builder

import (
	context "context"
	io "io"
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockDockerEngineClient is a mock of DockerEngineClient interface.
type MockDockerEngineClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerEngineClientMockRecorder
}

// MockDockerEngineClientMockRecorder is the mock recorder for MockDockerEngineClient.
type MockDockerEngineClientMockRecorder struct {
	mock *MockDockerEngineClient
}

// NewMockDockerEngineClient creates a new mock instance.
func NewMockDockerEngineClient(ctrl *gomock.Controller) *MockDockerEngineClient {
	mock := &MockDockerEngineClient{ctrl: ctrl}
	mock.recorder = &MockDockerEngineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerEngineClient) EXPECT() *MockDockerEngineClientMockRecorder {
	return m.recorder
}

// ContainerCreate mocks base method.
func (m *MockDockerEngineClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.ContainerCreateCreatedBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerCreate", ctx, config, hostConfig, networkingConfig, platform, containerName)
	ret0, _ := ret[0].(container.ContainerCreateCreatedBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerCreate indicates an expected call of ContainerCreate.
func (mr *MockDockerEngineClientMockRecorder) ContainerCreate(ctx, config, hostConfig, networkingConfig, platform, containerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerCreate", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerCreate), ctx, config, hostConfig, networkingConfig, platform, containerName)
}

// ContainerExecAttach mocks base method.
func (m *MockDockerEngineClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecAttach", ctx, execID, config)
	ret0, _ := ret[0].(types.HijackedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecAttach indicates an expected call of ContainerExecAttach.
func (mr *MockDockerEngineClientMockRecorder) ContainerExecAttach(ctx, execID, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecAttach", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerExecAttach), ctx, execID, config)
}

// ContainerExecCreate mocks base method.
func (m *MockDockerEngineClient) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecCreate", ctx, container, config)
	ret0, _ := ret[0].(types.IDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecCreate indicates an expected call of ContainerExecCreate.
func (mr *MockDockerEngineClientMockRecorder) ContainerExecCreate(ctx, container, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecCreate", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerExecCreate), ctx, container, config)
}

// ContainerExecInspect mocks base method.
func (m *MockDockerEngineClient) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecInspect", ctx, execID)
	ret0, _ := ret[0].(types.ContainerExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecInspect indicates an expected call of ContainerExecInspect.
func (mr *MockDockerEngineClientMockRecorder) ContainerExecInspect(ctx, execID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecInspect", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerExecInspect), ctx, execID)
}

// ContainerExecStart mocks base method.
func (m *MockDockerEngineClient) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecStart", ctx, execID, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerExecStart indicates an expected call of ContainerExecStart.
func (mr *MockDockerEngineClientMockRecorder) ContainerExecStart(ctx, execID, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecStart", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerExecStart), ctx, execID, config)
}

// ContainerLogs mocks base method.
func (m *MockDockerEngineClient) ContainerLogs(ctx context.Context, containerName string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerLogs", ctx, containerName, options)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs.
func (mr *MockDockerEngineClientMockRecorder) ContainerLogs(ctx, containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerLogs), ctx, containerName, options)
}

// ContainerRemove mocks base method.
func (m *MockDockerEngineClient) ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerRemove", ctx, container, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerRemove indicates an expected call of ContainerRemove.
func (mr *MockDockerEngineClientMockRecorder) ContainerRemove(ctx, container, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRemove", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerRemove), ctx, container, options)
}

// ContainerStart mocks base method.
func (m *MockDockerEngineClient) ContainerStart(ctx context.Context, containerName string, options types.ContainerStartOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerStart", ctx, containerName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStart indicates an expected call of ContainerStart.
func (mr *MockDockerEngineClientMockRecorder) ContainerStart(ctx, containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStart", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerStart), ctx, containerName, options)
}

// ContainerWait mocks base method.
func (m *MockDockerEngineClient) ContainerWait(ctx context.Context, containerName string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerWait", ctx, containerName, condition)
	ret0, _ := ret[0].(<-chan container.ContainerWaitOKBody)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// ContainerWait indicates an expected call of ContainerWait.
func (mr *MockDockerEngineClientMockRecorder) ContainerWait(ctx, containerName, condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerWait", reflect.TypeOf((*MockDockerEngineClient)(nil).ContainerWait), ctx, containerName, condition)
}

// CopyToContainer mocks base method.
func (m *MockDockerEngineClient) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyToContainer", ctx, container, path, content, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyToContainer indicates an expected call of CopyToContainer.
func (mr *MockDockerEngineClientMockRecorder) CopyToContainer(ctx, container, path, content, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyToContainer", reflect.TypeOf((*MockDockerEngineClient)(nil).CopyToContainer), ctx, container, path, content, options)
}

// ImageBuild mocks base method.
func (m *MockDockerEngineClient) ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageBuild", ctx, context, options)
	ret0, _ := ret[0].(types.ImageBuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBuild indicates an expected call of ImageBuild.
func (mr *MockDockerEngineClientMockRecorder) ImageBuild(ctx, context, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBuild", reflect.TypeOf((*MockDockerEngineClient)(nil).ImageBuild), ctx, context, options)
}

// ImageList mocks base method.
func (m *MockDockerEngineClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageList", ctx, options)
	ret0, _ := ret[0].([]types.ImageSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageList indicates an expected call of ImageList.
func (mr *MockDockerEngineClientMockRecorder) ImageList(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageList", reflect.TypeOf((*MockDockerEngineClient)(nil).ImageList), ctx, options)
}

// ImagePull mocks base method.
func (m *MockDockerEngineClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePull", ctx, refStr, options)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagePull indicates an expected call of ImagePull.
func (mr *MockDockerEngineClientMockRecorder) ImagePull(ctx, refStr, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePull", reflect.TypeOf((*MockDockerEngineClient)(nil).ImagePull), ctx, refStr, options)
}
