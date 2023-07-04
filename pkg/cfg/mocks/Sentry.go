// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	sentry "github.com/getsentry/sentry-go"
	mock "github.com/stretchr/testify/mock"
)

// Sentry is an autogenerated mock type for the Sentry type
type Sentry struct {
	mock.Mock
}

type Sentry_Expecter struct {
	mock *mock.Mock
}

func (_m *Sentry) EXPECT() *Sentry_Expecter {
	return &Sentry_Expecter{mock: &_m.Mock}
}

// CaptureException provides a mock function with given fields: exception, hint, scope
func (_m *Sentry) CaptureException(exception error, hint *sentry.EventHint, scope sentry.EventModifier) *sentry.EventID {
	ret := _m.Called(exception, hint, scope)

	var r0 *sentry.EventID
	if rf, ok := ret.Get(0).(func(error, *sentry.EventHint, sentry.EventModifier) *sentry.EventID); ok {
		r0 = rf(exception, hint, scope)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sentry.EventID)
		}
	}

	return r0
}

// Sentry_CaptureException_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CaptureException'
type Sentry_CaptureException_Call struct {
	*mock.Call
}

// CaptureException is a helper method to define mock.On call
//   - exception error
//   - hint *sentry.EventHint
//   - scope sentry.EventModifier
func (_e *Sentry_Expecter) CaptureException(exception interface{}, hint interface{}, scope interface{}) *Sentry_CaptureException_Call {
	return &Sentry_CaptureException_Call{Call: _e.mock.On("CaptureException", exception, hint, scope)}
}

func (_c *Sentry_CaptureException_Call) Run(run func(exception error, hint *sentry.EventHint, scope sentry.EventModifier)) *Sentry_CaptureException_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error), args[1].(*sentry.EventHint), args[2].(sentry.EventModifier))
	})
	return _c
}

func (_c *Sentry_CaptureException_Call) Return(_a0 *sentry.EventID) *Sentry_CaptureException_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sentry_CaptureException_Call) RunAndReturn(run func(error, *sentry.EventHint, sentry.EventModifier) *sentry.EventID) *Sentry_CaptureException_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSentry interface {
	mock.TestingT
	Cleanup(func())
}

// NewSentry creates a new instance of Sentry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSentry(t mockConstructorTestingTNewSentry) *Sentry {
	mock := &Sentry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
