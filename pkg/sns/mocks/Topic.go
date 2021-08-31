// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Topic is an autogenerated mock type for the Topic type
type Topic struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, msg, attributes
func (_m *Topic) Publish(ctx context.Context, msg *string, attributes ...map[string]interface{}) error {
	_va := make([]interface{}, len(attributes))
	for _i := range attributes {
		_va[_i] = attributes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, msg)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, ...map[string]interface{}) error); ok {
		r0 = rf(ctx, msg, attributes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeSqs provides a mock function with given fields: queueArn, attributes
func (_m *Topic) SubscribeSqs(queueArn string, attributes map[string]interface{}) error {
	ret := _m.Called(queueArn, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(queueArn, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
