// Code generated by MockGen. DO NOT EDIT.
// Source: ../clients/jsonplaceholder/client.go

// Package mockJsonplaceholderClient is a generated GoMock package.
package mockJsonplaceholderClient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jsonplaceholder "github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// FetchPost mocks base method.
func (m *MockClient) FetchPost(ctx context.Context, id string) (*jsonplaceholder.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPost", ctx, id)
	ret0, _ := ret[0].(*jsonplaceholder.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPost indicates an expected call of FetchPost.
func (mr *MockClientMockRecorder) FetchPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPost", reflect.TypeOf((*MockClient)(nil).FetchPost), ctx, id)
}

// FetchPosts mocks base method.
func (m *MockClient) FetchPosts(ctx context.Context) ([]*jsonplaceholder.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPosts", ctx)
	ret0, _ := ret[0].([]*jsonplaceholder.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPosts indicates an expected call of FetchPosts.
func (mr *MockClientMockRecorder) FetchPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPosts", reflect.TypeOf((*MockClient)(nil).FetchPosts), ctx)
}
