// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	dto "golang/module/transaction/dto"

	mock "github.com/stretchr/testify/mock"
)

// ControllerInterface is an autogenerated mock type for the ControllerInterface type
type ControllerInterface struct {
	mock.Mock
}

type ControllerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ControllerInterface) EXPECT() *ControllerInterface_Expecter {
	return &ControllerInterface_Expecter{mock: &_m.Mock}
}

// GetAllTransactionByRequest provides a mock function with given fields: req
func (_m *ControllerInterface) GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error, int64) {
	ret := _m.Called(req)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	var r2 int64
	if rf, ok := ret.Get(0).(func(*dto.Request) (*dto.GetAllResponseDataTransaction, error, int64)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*dto.Request) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(*dto.Request) int64); ok {
		r2 = rf(req)
	} else {
		r2 = ret.Get(2).(int64)
	}

	return r0, r1, r2
}

// ControllerInterface_GetAllTransactionByRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByRequest'
type ControllerInterface_GetAllTransactionByRequest_Call struct {
	*mock.Call
}

// GetAllTransactionByRequest is a helper method to define mock.On call
//   - req *dto.Request
func (_e *ControllerInterface_Expecter) GetAllTransactionByRequest(req interface{}) *ControllerInterface_GetAllTransactionByRequest_Call {
	return &ControllerInterface_GetAllTransactionByRequest_Call{Call: _e.mock.On("GetAllTransactionByRequest", req)}
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) Run(run func(req *dto.Request)) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dto.Request))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error, _a2 int64) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) RunAndReturn(run func(*dto.Request) (*dto.GetAllResponseDataTransaction, error, int64)) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Return(run)
	return _c
}

// NewControllerInterface creates a new instance of ControllerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewControllerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ControllerInterface {
	mock := &ControllerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}


type mockConstructorTestingTNewController interface {
	mock.TestingT
	Cleanup(func())
}

func NewController(t mockConstructorTestingTNewController) *ControllerInterface {
	mock := &ControllerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

