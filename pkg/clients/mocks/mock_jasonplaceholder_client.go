// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients/interface.go

// Package mockClients is a generated GoMock package.
package mockClients

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jsonplaceholder "github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
)

// MockJsonPlaceholderWriterClient is a mock of JsonPlaceholderWriterClient interface.
type MockJsonPlaceholderWriterClient struct {
	ctrl     *gomock.Controller
	recorder *MockJsonPlaceholderWriterClientMockRecorder
}

// MockJsonPlaceholderWriterClientMockRecorder is the mock recorder for MockJsonPlaceholderWriterClient.
type MockJsonPlaceholderWriterClientMockRecorder struct {
	mock *MockJsonPlaceholderWriterClient
}

// NewMockJsonPlaceholderWriterClient creates a new mock instance.
func NewMockJsonPlaceholderWriterClient(ctrl *gomock.Controller) *MockJsonPlaceholderWriterClient {
	mock := &MockJsonPlaceholderWriterClient{ctrl: ctrl}
	mock.recorder = &MockJsonPlaceholderWriterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJsonPlaceholderWriterClient) EXPECT() *MockJsonPlaceholderWriterClientMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockJsonPlaceholderWriterClient) Publish(post *jsonplaceholder.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockJsonPlaceholderWriterClientMockRecorder) Publish(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockJsonPlaceholderWriterClient)(nil).Publish), post)
}

// MockJsonPlaceholderClient is a mock of JsonPlaceholderClient interface.
type MockJsonPlaceholderClient struct {
	ctrl     *gomock.Controller
	recorder *MockJsonPlaceholderClientMockRecorder
}

// MockJsonPlaceholderClientMockRecorder is the mock recorder for MockJsonPlaceholderClient.
type MockJsonPlaceholderClientMockRecorder struct {
	mock *MockJsonPlaceholderClient
}

// NewMockJsonPlaceholderClient creates a new mock instance.
func NewMockJsonPlaceholderClient(ctrl *gomock.Controller) *MockJsonPlaceholderClient {
	mock := &MockJsonPlaceholderClient{ctrl: ctrl}
	mock.recorder = &MockJsonPlaceholderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJsonPlaceholderClient) EXPECT() *MockJsonPlaceholderClientMockRecorder {
	return m.recorder
}

// FetchPost mocks base method.
func (m *MockJsonPlaceholderClient) FetchPost(ctx context.Context, id string) (*jsonplaceholder.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPost", ctx, id)
	ret0, _ := ret[0].(*jsonplaceholder.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPost indicates an expected call of FetchPost.
func (mr *MockJsonPlaceholderClientMockRecorder) FetchPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPost", reflect.TypeOf((*MockJsonPlaceholderClient)(nil).FetchPost), ctx, id)
}

// FetchPosts mocks base method.
func (m *MockJsonPlaceholderClient) FetchPosts(ctx context.Context) ([]*jsonplaceholder.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPosts", ctx)
	ret0, _ := ret[0].([]*jsonplaceholder.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPosts indicates an expected call of FetchPosts.
func (mr *MockJsonPlaceholderClientMockRecorder) FetchPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPosts", reflect.TypeOf((*MockJsonPlaceholderClient)(nil).FetchPosts), ctx)
}