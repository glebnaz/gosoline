// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Topic is an autogenerated mock type for the Topic type
type Topic struct {
	mock.Mock
}

type Topic_Expecter struct {
	mock *mock.Mock
}

func (_m *Topic) EXPECT() *Topic_Expecter {
	return &Topic_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, msg, attributes
func (_m *Topic) Publish(ctx context.Context, msg string, attributes ...map[string]interface{}) error {
	_va := make([]interface{}, len(attributes))
	for _i := range attributes {
		_va[_i] = attributes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...map[string]interface{}) error); ok {
		r0 = rf(ctx, msg, attributes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topic_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type Topic_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - ctx context.Context
//   - msg string
//   - attributes ...map[string]interface{}
func (_e *Topic_Expecter) Publish(ctx interface{}, msg interface{}, attributes ...interface{}) *Topic_Publish_Call {
	return &Topic_Publish_Call{Call: _e.mock.On("Publish",
		append([]interface{}{ctx, msg}, attributes...)...)}
}

func (_c *Topic_Publish_Call) Run(run func(ctx context.Context, msg string, attributes ...map[string]interface{})) *Topic_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]map[string]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(map[string]interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Topic_Publish_Call) Return(_a0 error) *Topic_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Topic_Publish_Call) RunAndReturn(run func(context.Context, string, ...map[string]interface{}) error) *Topic_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// PublishBatch provides a mock function with given fields: ctx, messages, attributes
func (_m *Topic) PublishBatch(ctx context.Context, messages []string, attributes []map[string]interface{}) error {
	ret := _m.Called(ctx, messages, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, []map[string]interface{}) error); ok {
		r0 = rf(ctx, messages, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topic_PublishBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PublishBatch'
type Topic_PublishBatch_Call struct {
	*mock.Call
}

// PublishBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - messages []string
//   - attributes []map[string]interface{}
func (_e *Topic_Expecter) PublishBatch(ctx interface{}, messages interface{}, attributes interface{}) *Topic_PublishBatch_Call {
	return &Topic_PublishBatch_Call{Call: _e.mock.On("PublishBatch", ctx, messages, attributes)}
}

func (_c *Topic_PublishBatch_Call) Run(run func(ctx context.Context, messages []string, attributes []map[string]interface{})) *Topic_PublishBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].([]map[string]interface{}))
	})
	return _c
}

func (_c *Topic_PublishBatch_Call) Return(_a0 error) *Topic_PublishBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Topic_PublishBatch_Call) RunAndReturn(run func(context.Context, []string, []map[string]interface{}) error) *Topic_PublishBatch_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeSqs provides a mock function with given fields: ctx, queueArn, attributes
func (_m *Topic) SubscribeSqs(ctx context.Context, queueArn string, attributes map[string]interface{}) error {
	ret := _m.Called(ctx, queueArn, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) error); ok {
		r0 = rf(ctx, queueArn, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topic_SubscribeSqs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeSqs'
type Topic_SubscribeSqs_Call struct {
	*mock.Call
}

// SubscribeSqs is a helper method to define mock.On call
//   - ctx context.Context
//   - queueArn string
//   - attributes map[string]interface{}
func (_e *Topic_Expecter) SubscribeSqs(ctx interface{}, queueArn interface{}, attributes interface{}) *Topic_SubscribeSqs_Call {
	return &Topic_SubscribeSqs_Call{Call: _e.mock.On("SubscribeSqs", ctx, queueArn, attributes)}
}

func (_c *Topic_SubscribeSqs_Call) Run(run func(ctx context.Context, queueArn string, attributes map[string]interface{})) *Topic_SubscribeSqs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *Topic_SubscribeSqs_Call) Return(_a0 error) *Topic_SubscribeSqs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Topic_SubscribeSqs_Call) RunAndReturn(run func(context.Context, string, map[string]interface{}) error) *Topic_SubscribeSqs_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewTopic interface {
	mock.TestingT
	Cleanup(func())
}

// NewTopic creates a new instance of Topic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTopic(t mockConstructorTestingTNewTopic) *Topic {
	mock := &Topic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
