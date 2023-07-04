// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	blob "github.com/justtrackio/gosoline/pkg/blob"
	mock "github.com/stretchr/testify/mock"
)

// Stream is an autogenerated mock type for the Stream type
type Stream struct {
	mock.Mock
}

type Stream_Expecter struct {
	mock *mock.Mock
}

func (_m *Stream) EXPECT() *Stream_Expecter {
	return &Stream_Expecter{mock: &_m.Mock}
}

// AsReader provides a mock function with given fields:
func (_m *Stream) AsReader() blob.ReadSeekerCloser {
	ret := _m.Called()

	var r0 blob.ReadSeekerCloser
	if rf, ok := ret.Get(0).(func() blob.ReadSeekerCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blob.ReadSeekerCloser)
		}
	}

	return r0
}

// Stream_AsReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AsReader'
type Stream_AsReader_Call struct {
	*mock.Call
}

// AsReader is a helper method to define mock.On call
func (_e *Stream_Expecter) AsReader() *Stream_AsReader_Call {
	return &Stream_AsReader_Call{Call: _e.mock.On("AsReader")}
}

func (_c *Stream_AsReader_Call) Run(run func()) *Stream_AsReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Stream_AsReader_Call) Return(_a0 blob.ReadSeekerCloser) *Stream_AsReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Stream_AsReader_Call) RunAndReturn(run func() blob.ReadSeekerCloser) *Stream_AsReader_Call {
	_c.Call.Return(run)
	return _c
}

// ReadAll provides a mock function with given fields:
func (_m *Stream) ReadAll() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stream_ReadAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadAll'
type Stream_ReadAll_Call struct {
	*mock.Call
}

// ReadAll is a helper method to define mock.On call
func (_e *Stream_Expecter) ReadAll() *Stream_ReadAll_Call {
	return &Stream_ReadAll_Call{Call: _e.mock.On("ReadAll")}
}

func (_c *Stream_ReadAll_Call) Run(run func()) *Stream_ReadAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Stream_ReadAll_Call) Return(_a0 []byte, _a1 error) *Stream_ReadAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Stream_ReadAll_Call) RunAndReturn(run func() ([]byte, error)) *Stream_ReadAll_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStream interface {
	mock.TestingT
	Cleanup(func())
}

// NewStream creates a new instance of Stream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStream(t mockConstructorTestingTNewStream) *Stream {
	mock := &Stream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
