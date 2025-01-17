// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heavenopener/gohbase/hrpc (interfaces: RegionClient)

package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/heavenopener/gohbase/hrpc"
	reflect "reflect"
)

// MockRegionClient is a mock of RegionClient interface
type MockRegionClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegionClientMockRecorder
}

// MockRegionClientMockRecorder is the mock recorder for MockRegionClient
type MockRegionClientMockRecorder struct {
	mock *MockRegionClient
}

// NewMockRegionClient creates a new mock instance
func NewMockRegionClient(ctrl *gomock.Controller) *MockRegionClient {
	mock := &MockRegionClient{ctrl: ctrl}
	mock.recorder = &MockRegionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockRegionClient) EXPECT() *MockRegionClientMockRecorder {
	return _m.recorder
}

// Addr mocks base method
func (_m *MockRegionClient) Addr() string {
	ret := _m.ctrl.Call(_m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr
func (_mr *MockRegionClientMockRecorder) Addr() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Addr", reflect.TypeOf((*MockRegionClient)(nil).Addr))
}

// Close mocks base method
func (_m *MockRegionClient) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockRegionClientMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockRegionClient)(nil).Close))
}

// QueueRPC mocks base method
func (_m *MockRegionClient) QueueRPC(_param0 hrpc.Call) {
	_m.ctrl.Call(_m, "QueueRPC", _param0)
}

// QueueRPC indicates an expected call of QueueRPC
func (_mr *MockRegionClientMockRecorder) QueueRPC(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueueRPC", reflect.TypeOf((*MockRegionClient)(nil).QueueRPC), arg0)
}

// String mocks base method
func (_m *MockRegionClient) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (_mr *MockRegionClientMockRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "String", reflect.TypeOf((*MockRegionClient)(nil).String))
}
