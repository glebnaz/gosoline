// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	parquet "github.com/justtrackio/gosoline/pkg/parquet"
	mock "github.com/stretchr/testify/mock"
)

// FileRecorder is an autogenerated mock type for the FileRecorder type
type FileRecorder struct {
	mock.Mock
}

type FileRecorder_Expecter struct {
	mock *mock.Mock
}

func (_m *FileRecorder) EXPECT() *FileRecorder_Expecter {
	return &FileRecorder_Expecter{mock: &_m.Mock}
}

// DeleteRecordedFiles provides a mock function with given fields: ctx
func (_m *FileRecorder) DeleteRecordedFiles(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileRecorder_DeleteRecordedFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRecordedFiles'
type FileRecorder_DeleteRecordedFiles_Call struct {
	*mock.Call
}

// DeleteRecordedFiles is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FileRecorder_Expecter) DeleteRecordedFiles(ctx interface{}) *FileRecorder_DeleteRecordedFiles_Call {
	return &FileRecorder_DeleteRecordedFiles_Call{Call: _e.mock.On("DeleteRecordedFiles", ctx)}
}

func (_c *FileRecorder_DeleteRecordedFiles_Call) Run(run func(ctx context.Context)) *FileRecorder_DeleteRecordedFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FileRecorder_DeleteRecordedFiles_Call) Return(_a0 error) *FileRecorder_DeleteRecordedFiles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileRecorder_DeleteRecordedFiles_Call) RunAndReturn(run func(context.Context) error) *FileRecorder_DeleteRecordedFiles_Call {
	_c.Call.Return(run)
	return _c
}

// Files provides a mock function with given fields:
func (_m *FileRecorder) Files() []parquet.File {
	ret := _m.Called()

	var r0 []parquet.File
	if rf, ok := ret.Get(0).(func() []parquet.File); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]parquet.File)
		}
	}

	return r0
}

// FileRecorder_Files_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Files'
type FileRecorder_Files_Call struct {
	*mock.Call
}

// Files is a helper method to define mock.On call
func (_e *FileRecorder_Expecter) Files() *FileRecorder_Files_Call {
	return &FileRecorder_Files_Call{Call: _e.mock.On("Files")}
}

func (_c *FileRecorder_Files_Call) Run(run func()) *FileRecorder_Files_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileRecorder_Files_Call) Return(_a0 []parquet.File) *FileRecorder_Files_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileRecorder_Files_Call) RunAndReturn(run func() []parquet.File) *FileRecorder_Files_Call {
	_c.Call.Return(run)
	return _c
}

// RecordFile provides a mock function with given fields: bucket, key
func (_m *FileRecorder) RecordFile(bucket string, key string) {
	_m.Called(bucket, key)
}

// FileRecorder_RecordFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecordFile'
type FileRecorder_RecordFile_Call struct {
	*mock.Call
}

// RecordFile is a helper method to define mock.On call
//   - bucket string
//   - key string
func (_e *FileRecorder_Expecter) RecordFile(bucket interface{}, key interface{}) *FileRecorder_RecordFile_Call {
	return &FileRecorder_RecordFile_Call{Call: _e.mock.On("RecordFile", bucket, key)}
}

func (_c *FileRecorder_RecordFile_Call) Run(run func(bucket string, key string)) *FileRecorder_RecordFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *FileRecorder_RecordFile_Call) Return() *FileRecorder_RecordFile_Call {
	_c.Call.Return()
	return _c
}

func (_c *FileRecorder_RecordFile_Call) RunAndReturn(run func(string, string)) *FileRecorder_RecordFile_Call {
	_c.Call.Return(run)
	return _c
}

// RenameRecordedFiles provides a mock function with given fields: ctx, newPrefix
func (_m *FileRecorder) RenameRecordedFiles(ctx context.Context, newPrefix string) error {
	ret := _m.Called(ctx, newPrefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, newPrefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileRecorder_RenameRecordedFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenameRecordedFiles'
type FileRecorder_RenameRecordedFiles_Call struct {
	*mock.Call
}

// RenameRecordedFiles is a helper method to define mock.On call
//   - ctx context.Context
//   - newPrefix string
func (_e *FileRecorder_Expecter) RenameRecordedFiles(ctx interface{}, newPrefix interface{}) *FileRecorder_RenameRecordedFiles_Call {
	return &FileRecorder_RenameRecordedFiles_Call{Call: _e.mock.On("RenameRecordedFiles", ctx, newPrefix)}
}

func (_c *FileRecorder_RenameRecordedFiles_Call) Run(run func(ctx context.Context, newPrefix string)) *FileRecorder_RenameRecordedFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileRecorder_RenameRecordedFiles_Call) Return(_a0 error) *FileRecorder_RenameRecordedFiles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileRecorder_RenameRecordedFiles_Call) RunAndReturn(run func(context.Context, string) error) *FileRecorder_RenameRecordedFiles_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFileRecorder interface {
	mock.TestingT
	Cleanup(func())
}

// NewFileRecorder creates a new instance of FileRecorder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFileRecorder(t mockConstructorTestingTNewFileRecorder) *FileRecorder {
	mock := &FileRecorder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
