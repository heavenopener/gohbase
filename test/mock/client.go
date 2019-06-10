// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heavenopener/gohbase (interfaces: Client)

package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/heavenopener/gohbase/hrpc"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClient) EXPECT() *MockClientMockRecorder {
	return _m.recorder
}

// Append mocks base method
func (_m *MockClient) Append(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Append", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Append indicates an expected call of Append
func (_mr *MockClientMockRecorder) Append(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Append", reflect.TypeOf((*MockClient)(nil).Append), arg0)
}

// CheckAndPut mocks base method
func (_m *MockClient) CheckAndPut(_param0 *hrpc.Mutate, _param1 string, _param2 string, _param3 []byte) (bool, error) {
	ret := _m.ctrl.Call(_m, "CheckAndPut", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndPut indicates an expected call of CheckAndPut
func (_mr *MockClientMockRecorder) CheckAndPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckAndPut", reflect.TypeOf((*MockClient)(nil).CheckAndPut), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (_m *MockClient) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockClientMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Delete mocks base method
func (_m *MockClient) Delete(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (_mr *MockClientMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0)
}

// Get mocks base method
func (_m *MockClient) Get(_param0 *hrpc.Get) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0)
}

// Increment mocks base method
func (_m *MockClient) Increment(_param0 *hrpc.Mutate) (int64, error) {
	ret := _m.ctrl.Call(_m, "Increment", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment
func (_mr *MockClientMockRecorder) Increment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Increment", reflect.TypeOf((*MockClient)(nil).Increment), arg0)
}

// Put mocks base method
func (_m *MockClient) Put(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Put", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (_mr *MockClientMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0)
}

// Scan mocks base method
func (_m *MockClient) Scan(_param0 *hrpc.Scan) hrpc.Scanner {
	ret := _m.ctrl.Call(_m, "Scan", _param0)
	ret0, _ := ret[0].(hrpc.Scanner)
	return ret0
}

// Scan indicates an expected call of Scan
func (_mr *MockClientMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Scan", reflect.TypeOf((*MockClient)(nil).Scan), arg0)
}
