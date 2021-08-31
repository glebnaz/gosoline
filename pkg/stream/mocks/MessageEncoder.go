// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	stream "github.com/applike/gosoline/pkg/stream"
	mock "github.com/stretchr/testify/mock"
)

// MessageEncoder is an autogenerated mock type for the MessageEncoder type
type MessageEncoder struct {
	mock.Mock
}

// Decode provides a mock function with given fields: ctx, msg, out
func (_m *MessageEncoder) Decode(ctx context.Context, msg *stream.Message, out interface{}) (context.Context, map[string]interface{}, error) {
	ret := _m.Called(ctx, msg, out)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, *stream.Message, interface{}) context.Context); ok {
		r0 = rf(ctx, msg, out)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(context.Context, *stream.Message, interface{}) map[string]interface{}); ok {
		r1 = rf(ctx, msg, out)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *stream.Message, interface{}) error); ok {
		r2 = rf(ctx, msg, out)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Encode provides a mock function with given fields: ctx, data, attributeSets
func (_m *MessageEncoder) Encode(ctx context.Context, data interface{}, attributeSets ...map[string]interface{}) (*stream.Message, error) {
	_va := make([]interface{}, len(attributeSets))
	for _i := range attributeSets {
		_va[_i] = attributeSets[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stream.Message
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...map[string]interface{}) *stream.Message); ok {
		r0 = rf(ctx, data, attributeSets...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stream.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...map[string]interface{}) error); ok {
		r1 = rf(ctx, data, attributeSets...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
