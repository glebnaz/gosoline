// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	servicesqs "github.com/aws/aws-sdk-go/service/sqs"
	mock "github.com/stretchr/testify/mock"

	sqs "github.com/applike/gosoline/pkg/sqs"
)

// Queue is an autogenerated mock type for the Queue type
type Queue struct {
	mock.Mock
}

// DeleteMessage provides a mock function with given fields: receiptHandle
func (_m *Queue) DeleteMessage(receiptHandle string) error {
	ret := _m.Called(receiptHandle)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(receiptHandle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessageBatch provides a mock function with given fields: receiptHandles
func (_m *Queue) DeleteMessageBatch(receiptHandles []string) error {
	ret := _m.Called(receiptHandles)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(receiptHandles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArn provides a mock function with given fields:
func (_m *Queue) GetArn() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *Queue) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUrl provides a mock function with given fields:
func (_m *Queue) GetUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Receive provides a mock function with given fields: ctx, maxNumberOfMessages, waitTime
func (_m *Queue) Receive(ctx context.Context, maxNumberOfMessages int64, waitTime int64) ([]*servicesqs.Message, error) {
	ret := _m.Called(ctx, maxNumberOfMessages, waitTime)

	var r0 []*servicesqs.Message
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*servicesqs.Message); ok {
		r0 = rf(ctx, maxNumberOfMessages, waitTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*servicesqs.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, maxNumberOfMessages, waitTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: ctx, msg
func (_m *Queue) Send(ctx context.Context, msg *sqs.Message) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqs.Message) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendBatch provides a mock function with given fields: ctx, messages
func (_m *Queue) SendBatch(ctx context.Context, messages []*sqs.Message) error {
	ret := _m.Called(ctx, messages)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*sqs.Message) error); ok {
		r0 = rf(ctx, messages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
