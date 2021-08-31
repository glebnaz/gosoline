// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/applike/gosoline/pkg/log"
	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: format, args
func (_m *Logger) Debug(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: format, args
func (_m *Logger) Error(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: format, args
func (_m *Logger) Info(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: format, args
func (_m *Logger) Warn(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithChannel provides a mock function with given fields: channel
func (_m *Logger) WithChannel(channel string) log.Logger {
	ret := _m.Called(channel)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(string) log.Logger); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Logger) WithContext(ctx context.Context) log.Logger {
	ret := _m.Called(ctx)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(context.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: _a0
func (_m *Logger) WithFields(_a0 log.Fields) log.Logger {
	ret := _m.Called(_a0)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(log.Fields) log.Logger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}
