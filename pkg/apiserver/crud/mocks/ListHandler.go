// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	crud "github.com/justtrackio/gosoline/pkg/apiserver/crud"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// ListHandler is an autogenerated mock type for the ListHandler type
type ListHandler struct {
	mock.Mock
}

type ListHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *ListHandler) EXPECT() *ListHandler_Expecter {
	return &ListHandler_Expecter{mock: &_m.Mock}
}

// GetModel provides a mock function with given fields:
func (_m *ListHandler) GetModel() db_repo.ModelBased {
	ret := _m.Called()

	var r0 db_repo.ModelBased
	if rf, ok := ret.Get(0).(func() db_repo.ModelBased); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db_repo.ModelBased)
		}
	}

	return r0
}

// ListHandler_GetModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModel'
type ListHandler_GetModel_Call struct {
	*mock.Call
}

// GetModel is a helper method to define mock.On call
func (_e *ListHandler_Expecter) GetModel() *ListHandler_GetModel_Call {
	return &ListHandler_GetModel_Call{Call: _e.mock.On("GetModel")}
}

func (_c *ListHandler_GetModel_Call) Run(run func()) *ListHandler_GetModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ListHandler_GetModel_Call) Return(_a0 db_repo.ModelBased) *ListHandler_GetModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ListHandler_GetModel_Call) RunAndReturn(run func() db_repo.ModelBased) *ListHandler_GetModel_Call {
	_c.Call.Return(run)
	return _c
}

// GetRepository provides a mock function with given fields:
func (_m *ListHandler) GetRepository() crud.Repository {
	ret := _m.Called()

	var r0 crud.Repository
	if rf, ok := ret.Get(0).(func() crud.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crud.Repository)
		}
	}

	return r0
}

// ListHandler_GetRepository_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRepository'
type ListHandler_GetRepository_Call struct {
	*mock.Call
}

// GetRepository is a helper method to define mock.On call
func (_e *ListHandler_Expecter) GetRepository() *ListHandler_GetRepository_Call {
	return &ListHandler_GetRepository_Call{Call: _e.mock.On("GetRepository")}
}

func (_c *ListHandler_GetRepository_Call) Run(run func()) *ListHandler_GetRepository_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ListHandler_GetRepository_Call) Return(_a0 crud.Repository) *ListHandler_GetRepository_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ListHandler_GetRepository_Call) RunAndReturn(run func() crud.Repository) *ListHandler_GetRepository_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, qb, apiView
func (_m *ListHandler) List(ctx context.Context, qb *db_repo.QueryBuilder, apiView string) (interface{}, error) {
	ret := _m.Called(ctx, qb, apiView)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, string) (interface{}, error)); ok {
		return rf(ctx, qb, apiView)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, string) interface{}); ok {
		r0 = rf(ctx, qb, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *db_repo.QueryBuilder, string) error); ok {
		r1 = rf(ctx, qb, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHandler_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ListHandler_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - apiView string
func (_e *ListHandler_Expecter) List(ctx interface{}, qb interface{}, apiView interface{}) *ListHandler_List_Call {
	return &ListHandler_List_Call{Call: _e.mock.On("List", ctx, qb, apiView)}
}

func (_c *ListHandler_List_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, apiView string)) *ListHandler_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(string))
	})
	return _c
}

func (_c *ListHandler_List_Call) Return(out interface{}, err error) *ListHandler_List_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *ListHandler_List_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, string) (interface{}, error)) *ListHandler_List_Call {
	_c.Call.Return(run)
	return _c
}

// TransformOutput provides a mock function with given fields: ctx, model, apiView
func (_m *ListHandler) TransformOutput(ctx context.Context, model db_repo.ModelBased, apiView string) (interface{}, error) {
	ret := _m.Called(ctx, model, apiView)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased, string) (interface{}, error)); ok {
		return rf(ctx, model, apiView)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased, string) interface{}); ok {
		r0 = rf(ctx, model, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db_repo.ModelBased, string) error); ok {
		r1 = rf(ctx, model, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHandler_TransformOutput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformOutput'
type ListHandler_TransformOutput_Call struct {
	*mock.Call
}

// TransformOutput is a helper method to define mock.On call
//   - ctx context.Context
//   - model db_repo.ModelBased
//   - apiView string
func (_e *ListHandler_Expecter) TransformOutput(ctx interface{}, model interface{}, apiView interface{}) *ListHandler_TransformOutput_Call {
	return &ListHandler_TransformOutput_Call{Call: _e.mock.On("TransformOutput", ctx, model, apiView)}
}

func (_c *ListHandler_TransformOutput_Call) Run(run func(ctx context.Context, model db_repo.ModelBased, apiView string)) *ListHandler_TransformOutput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_repo.ModelBased), args[2].(string))
	})
	return _c
}

func (_c *ListHandler_TransformOutput_Call) Return(output interface{}, err error) *ListHandler_TransformOutput_Call {
	_c.Call.Return(output, err)
	return _c
}

func (_c *ListHandler_TransformOutput_Call) RunAndReturn(run func(context.Context, db_repo.ModelBased, string) (interface{}, error)) *ListHandler_TransformOutput_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewListHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewListHandler creates a new instance of ListHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListHandler(t mockConstructorTestingTNewListHandler) *ListHandler {
	mock := &ListHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
