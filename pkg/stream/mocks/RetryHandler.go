// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	stream "github.com/justtrackio/gosoline/pkg/stream"
	mock "github.com/stretchr/testify/mock"
)

// RetryHandler is an autogenerated mock type for the RetryHandler type
type RetryHandler struct {
	mock.Mock
}

type RetryHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *RetryHandler) EXPECT() *RetryHandler_Expecter {
	return &RetryHandler_Expecter{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *RetryHandler) Data() <-chan *stream.Message {
	ret := _m.Called()

	var r0 <-chan *stream.Message
	if rf, ok := ret.Get(0).(func() <-chan *stream.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *stream.Message)
		}
	}

	return r0
}

// RetryHandler_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type RetryHandler_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *RetryHandler_Expecter) Data() *RetryHandler_Data_Call {
	return &RetryHandler_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *RetryHandler_Data_Call) Run(run func()) *RetryHandler_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RetryHandler_Data_Call) Return(_a0 <-chan *stream.Message) *RetryHandler_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RetryHandler_Data_Call) RunAndReturn(run func() <-chan *stream.Message) *RetryHandler_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, msg
func (_m *RetryHandler) Put(ctx context.Context, msg *stream.Message) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stream.Message) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetryHandler_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type RetryHandler_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *stream.Message
func (_e *RetryHandler_Expecter) Put(ctx interface{}, msg interface{}) *RetryHandler_Put_Call {
	return &RetryHandler_Put_Call{Call: _e.mock.On("Put", ctx, msg)}
}

func (_c *RetryHandler_Put_Call) Run(run func(ctx context.Context, msg *stream.Message)) *RetryHandler_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*stream.Message))
	})
	return _c
}

func (_c *RetryHandler_Put_Call) Return(_a0 error) *RetryHandler_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RetryHandler_Put_Call) RunAndReturn(run func(context.Context, *stream.Message) error) *RetryHandler_Put_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *RetryHandler) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetryHandler_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type RetryHandler_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RetryHandler_Expecter) Run(ctx interface{}) *RetryHandler_Run_Call {
	return &RetryHandler_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *RetryHandler_Run_Call) Run(run func(ctx context.Context)) *RetryHandler_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RetryHandler_Run_Call) Return(_a0 error) *RetryHandler_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RetryHandler_Run_Call) RunAndReturn(run func(context.Context) error) *RetryHandler_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *RetryHandler) Stop() {
	_m.Called()
}

// RetryHandler_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type RetryHandler_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *RetryHandler_Expecter) Stop() *RetryHandler_Stop_Call {
	return &RetryHandler_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *RetryHandler_Stop_Call) Run(run func()) *RetryHandler_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RetryHandler_Stop_Call) Return() *RetryHandler_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *RetryHandler_Stop_Call) RunAndReturn(run func()) *RetryHandler_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRetryHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewRetryHandler creates a new instance of RetryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRetryHandler(t mockConstructorTestingTNewRetryHandler) *RetryHandler {
	mock := &RetryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
