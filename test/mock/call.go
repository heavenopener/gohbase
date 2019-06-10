// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heavenopener/gohbase/hrpc (interfaces: Call)

package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	hrpc "github.com/heavenopener/gohbase/hrpc"
)

// MockCall is a mock of Call interface
type MockCall struct {
	ctrl     *gomock.Controller
	recorder *MockCallMockRecorder
}

// MockCallMockRecorder is the mock recorder for MockCall
type MockCallMockRecorder struct {
	mock *MockCall
}

// NewMockCall creates a new mock instance
func NewMockCall(ctrl *gomock.Controller) *MockCall {
	mock := &MockCall{ctrl: ctrl}
	mock.recorder = &MockCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCall) EXPECT() *MockCallMockRecorder {
	return _m.recorder
}

// Context mocks base method
func (_m *MockCall) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockCallMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockCall)(nil).Context))
}

// Key mocks base method
func (_m *MockCall) Key() []byte {
	ret := _m.ctrl.Call(_m, "Key")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Key indicates an expected call of Key
func (_mr *MockCallMockRecorder) Key() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Key", reflect.TypeOf((*MockCall)(nil).Key))
}

// Name mocks base method
func (_m *MockCall) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (_mr *MockCallMockRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Name", reflect.TypeOf((*MockCall)(nil).Name))
}

// NewResponse mocks base method
func (_m *MockCall) NewResponse() proto.Message {
	ret := _m.ctrl.Call(_m, "NewResponse")
	ret0, _ := ret[0].(proto.Message)
	return ret0
}

// NewResponse indicates an expected call of NewResponse
func (_mr *MockCallMockRecorder) NewResponse() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NewResponse", reflect.TypeOf((*MockCall)(nil).NewResponse))
}

// Region mocks base method
func (_m *MockCall) Region() hrpc.RegionInfo {
	ret := _m.ctrl.Call(_m, "Region")
	ret0, _ := ret[0].(hrpc.RegionInfo)
	return ret0
}

// Region indicates an expected call of Region
func (_mr *MockCallMockRecorder) Region() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Region", reflect.TypeOf((*MockCall)(nil).Region))
}

// ResultChan mocks base method
func (_m *MockCall) ResultChan() chan hrpc.RPCResult {
	ret := _m.ctrl.Call(_m, "ResultChan")
	ret0, _ := ret[0].(chan hrpc.RPCResult)
	return ret0
}

// ResultChan indicates an expected call of ResultChan
func (_mr *MockCallMockRecorder) ResultChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ResultChan", reflect.TypeOf((*MockCall)(nil).ResultChan))
}

// SetRegion mocks base method
func (_m *MockCall) SetRegion(_param0 hrpc.RegionInfo) {
	_m.ctrl.Call(_m, "SetRegion", _param0)
}

// SetRegion indicates an expected call of SetRegion
func (_mr *MockCallMockRecorder) SetRegion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRegion", reflect.TypeOf((*MockCall)(nil).SetRegion), arg0)
}

// Table mocks base method
func (_m *MockCall) Table() []byte {
	ret := _m.ctrl.Call(_m, "Table")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Table indicates an expected call of Table
func (_mr *MockCallMockRecorder) Table() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Table", reflect.TypeOf((*MockCall)(nil).Table))
}

// ToProto mocks base method
func (_m *MockCall) ToProto() proto.Message {
	ret := _m.ctrl.Call(_m, "ToProto")
	ret0, _ := ret[0].(proto.Message)
	return ret0
}

// ToProto indicates an expected call of ToProto
func (_mr *MockCallMockRecorder) ToProto() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ToProto", reflect.TypeOf((*MockCall)(nil).ToProto))
}
