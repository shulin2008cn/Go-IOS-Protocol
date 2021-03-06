// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/Go-IOS-Protocol/rpc (interfaces: CliServer)

// Package rpc_mock is a generated GoMock package.
package rpc_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rpc "github.com/iost-official/Go-IOS-Protocol/rpc"
)

// MockCliServer is a mock of CliServer interface
type MockCliServer struct {
	ctrl     *gomock.Controller
	recorder *MockCliServerMockRecorder
}

// MockCliServerMockRecorder is the mock recorder for MockCliServer
type MockCliServerMockRecorder struct {
	mock *MockCliServer
}

// NewMockCliServer creates a new mock instance
func NewMockCliServer(ctrl *gomock.Controller) *MockCliServer {
	mock := &MockCliServer{ctrl: ctrl}
	mock.recorder = &MockCliServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliServer) EXPECT() *MockCliServerMockRecorder {
	return m.recorder
}

// GetBalance mocks base method
func (m *MockCliServer) GetBalance(arg0 context.Context, arg1 *rpc.Key) (*rpc.Value, error) {
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1)
	ret0, _ := ret[0].(*rpc.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockCliServerMockRecorder) GetBalance(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockCliServer)(nil).GetBalance), arg0, arg1)
}

// GetBlock mocks base method
func (m *MockCliServer) GetBlock(arg0 context.Context, arg1 *rpc.BlockKey) (*rpc.BlockInfo, error) {
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*rpc.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockCliServerMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockCliServer)(nil).GetBlock), arg0, arg1)
}

// GetBlockByHeight mocks base method
func (m *MockCliServer) GetBlockByHeight(arg0 context.Context, arg1 *rpc.BlockKey) (*rpc.BlockInfo, error) {
	ret := m.ctrl.Call(m, "GetBlockByHeight", arg0, arg1)
	ret0, _ := ret[0].(*rpc.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight
func (mr *MockCliServerMockRecorder) GetBlockByHeight(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockCliServer)(nil).GetBlockByHeight), arg0, arg1)
}

// GetState mocks base method
func (m *MockCliServer) GetState(arg0 context.Context, arg1 *rpc.Key) (*rpc.Value, error) {
	ret := m.ctrl.Call(m, "GetState", arg0, arg1)
	ret0, _ := ret[0].(*rpc.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState
func (mr *MockCliServerMockRecorder) GetState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockCliServer)(nil).GetState), arg0, arg1)
}

// GetTransaction mocks base method
func (m *MockCliServer) GetTransaction(arg0 context.Context, arg1 *rpc.TransactionKey) (*rpc.Transaction, error) {
	ret := m.ctrl.Call(m, "GetTransaction", arg0, arg1)
	ret0, _ := ret[0].(*rpc.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockCliServerMockRecorder) GetTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockCliServer)(nil).GetTransaction), arg0, arg1)
}

// GetTransactionByHash mocks base method
func (m *MockCliServer) GetTransactionByHash(arg0 context.Context, arg1 *rpc.TransactionHash) (*rpc.Transaction, error) {
	ret := m.ctrl.Call(m, "GetTransactionByHash", arg0, arg1)
	ret0, _ := ret[0].(*rpc.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash
func (mr *MockCliServerMockRecorder) GetTransactionByHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockCliServer)(nil).GetTransactionByHash), arg0, arg1)
}

// PublishTx mocks base method
func (m *MockCliServer) PublishTx(arg0 context.Context, arg1 *rpc.Transaction) (*rpc.Response, error) {
	ret := m.ctrl.Call(m, "PublishTx", arg0, arg1)
	ret0, _ := ret[0].(*rpc.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishTx indicates an expected call of PublishTx
func (mr *MockCliServerMockRecorder) PublishTx(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTx", reflect.TypeOf((*MockCliServer)(nil).PublishTx), arg0, arg1)
}
