// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KvStore is an autogenerated mock type for the KvStore type
type KvStore[T interface{}] struct {
	mock.Mock
}

type KvStore_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *KvStore[T]) EXPECT() *KvStore_Expecter[T] {
	return &KvStore_Expecter[T]{mock: &_m.Mock}
}

// Contains provides a mock function with given fields: ctx, key
func (_m *KvStore[T]) Contains(ctx context.Context, key interface{}) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (bool, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KvStore_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type KvStore_Contains_Call[T interface{}] struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - ctx context.Context
//   - key interface{}
func (_e *KvStore_Expecter[T]) Contains(ctx interface{}, key interface{}) *KvStore_Contains_Call[T] {
	return &KvStore_Contains_Call[T]{Call: _e.mock.On("Contains", ctx, key)}
}

func (_c *KvStore_Contains_Call[T]) Run(run func(ctx context.Context, key interface{})) *KvStore_Contains_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *KvStore_Contains_Call[T]) Return(_a0 bool, _a1 error) *KvStore_Contains_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KvStore_Contains_Call[T]) RunAndReturn(run func(context.Context, interface{}) (bool, error)) *KvStore_Contains_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, key
func (_m *KvStore[T]) Delete(ctx context.Context, key interface{}) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KvStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type KvStore_Delete_Call[T interface{}] struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key interface{}
func (_e *KvStore_Expecter[T]) Delete(ctx interface{}, key interface{}) *KvStore_Delete_Call[T] {
	return &KvStore_Delete_Call[T]{Call: _e.mock.On("Delete", ctx, key)}
}

func (_c *KvStore_Delete_Call[T]) Run(run func(ctx context.Context, key interface{})) *KvStore_Delete_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *KvStore_Delete_Call[T]) Return(_a0 error) *KvStore_Delete_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KvStore_Delete_Call[T]) RunAndReturn(run func(context.Context, interface{}) error) *KvStore_Delete_Call[T] {
	_c.Call.Return(run)
	return _c
}

// DeleteBatch provides a mock function with given fields: ctx, keys
func (_m *KvStore[T]) DeleteBatch(ctx context.Context, keys interface{}) error {
	ret := _m.Called(ctx, keys)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KvStore_DeleteBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBatch'
type KvStore_DeleteBatch_Call[T interface{}] struct {
	*mock.Call
}

// DeleteBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - keys interface{}
func (_e *KvStore_Expecter[T]) DeleteBatch(ctx interface{}, keys interface{}) *KvStore_DeleteBatch_Call[T] {
	return &KvStore_DeleteBatch_Call[T]{Call: _e.mock.On("DeleteBatch", ctx, keys)}
}

func (_c *KvStore_DeleteBatch_Call[T]) Run(run func(ctx context.Context, keys interface{})) *KvStore_DeleteBatch_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *KvStore_DeleteBatch_Call[T]) Return(_a0 error) *KvStore_DeleteBatch_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KvStore_DeleteBatch_Call[T]) RunAndReturn(run func(context.Context, interface{}) error) *KvStore_DeleteBatch_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *KvStore[T]) Get(ctx context.Context, key interface{}, value *T) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *T) (bool, error)); ok {
		return rf(ctx, key, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *T) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, *T) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KvStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type KvStore_Get_Call[T interface{}] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key interface{}
//   - value *T
func (_e *KvStore_Expecter[T]) Get(ctx interface{}, key interface{}, value interface{}) *KvStore_Get_Call[T] {
	return &KvStore_Get_Call[T]{Call: _e.mock.On("Get", ctx, key, value)}
}

func (_c *KvStore_Get_Call[T]) Run(run func(ctx context.Context, key interface{}, value *T)) *KvStore_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(*T))
	})
	return _c
}

func (_c *KvStore_Get_Call[T]) Return(_a0 bool, _a1 error) *KvStore_Get_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KvStore_Get_Call[T]) RunAndReturn(run func(context.Context, interface{}, *T) (bool, error)) *KvStore_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetBatch provides a mock function with given fields: ctx, keys, values
func (_m *KvStore[T]) GetBatch(ctx context.Context, keys interface{}, values interface{}) ([]interface{}, error) {
	ret := _m.Called(ctx, keys, values)

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) ([]interface{}, error)); ok {
		return rf(ctx, keys, values)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) []interface{}); ok {
		r0 = rf(ctx, keys, values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}) error); ok {
		r1 = rf(ctx, keys, values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KvStore_GetBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBatch'
type KvStore_GetBatch_Call[T interface{}] struct {
	*mock.Call
}

// GetBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - keys interface{}
//   - values interface{}
func (_e *KvStore_Expecter[T]) GetBatch(ctx interface{}, keys interface{}, values interface{}) *KvStore_GetBatch_Call[T] {
	return &KvStore_GetBatch_Call[T]{Call: _e.mock.On("GetBatch", ctx, keys, values)}
}

func (_c *KvStore_GetBatch_Call[T]) Run(run func(ctx context.Context, keys interface{}, values interface{})) *KvStore_GetBatch_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(interface{}))
	})
	return _c
}

func (_c *KvStore_GetBatch_Call[T]) Return(_a0 []interface{}, _a1 error) *KvStore_GetBatch_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KvStore_GetBatch_Call[T]) RunAndReturn(run func(context.Context, interface{}, interface{}) ([]interface{}, error)) *KvStore_GetBatch_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, key, value
func (_m *KvStore[T]) Put(ctx context.Context, key interface{}, value T) error {
	ret := _m.Called(ctx, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, T) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KvStore_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type KvStore_Put_Call[T interface{}] struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - key interface{}
//   - value T
func (_e *KvStore_Expecter[T]) Put(ctx interface{}, key interface{}, value interface{}) *KvStore_Put_Call[T] {
	return &KvStore_Put_Call[T]{Call: _e.mock.On("Put", ctx, key, value)}
}

func (_c *KvStore_Put_Call[T]) Run(run func(ctx context.Context, key interface{}, value T)) *KvStore_Put_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(T))
	})
	return _c
}

func (_c *KvStore_Put_Call[T]) Return(_a0 error) *KvStore_Put_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KvStore_Put_Call[T]) RunAndReturn(run func(context.Context, interface{}, T) error) *KvStore_Put_Call[T] {
	_c.Call.Return(run)
	return _c
}

// PutBatch provides a mock function with given fields: ctx, values
func (_m *KvStore[T]) PutBatch(ctx context.Context, values interface{}) error {
	ret := _m.Called(ctx, values)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KvStore_PutBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutBatch'
type KvStore_PutBatch_Call[T interface{}] struct {
	*mock.Call
}

// PutBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - values interface{}
func (_e *KvStore_Expecter[T]) PutBatch(ctx interface{}, values interface{}) *KvStore_PutBatch_Call[T] {
	return &KvStore_PutBatch_Call[T]{Call: _e.mock.On("PutBatch", ctx, values)}
}

func (_c *KvStore_PutBatch_Call[T]) Run(run func(ctx context.Context, values interface{})) *KvStore_PutBatch_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *KvStore_PutBatch_Call[T]) Return(_a0 error) *KvStore_PutBatch_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *KvStore_PutBatch_Call[T]) RunAndReturn(run func(context.Context, interface{}) error) *KvStore_PutBatch_Call[T] {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewKvStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewKvStore creates a new instance of KvStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKvStore[T interface{}](t mockConstructorTestingTNewKvStore) *KvStore[T] {
	mock := &KvStore[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
