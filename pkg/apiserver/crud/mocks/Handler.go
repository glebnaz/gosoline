// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	crud "github.com/justtrackio/gosoline/pkg/apiserver/crud"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// GetCreateInput provides a mock function with given fields:
func (_m *Handler) GetCreateInput() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Handler_GetCreateInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreateInput'
type Handler_GetCreateInput_Call struct {
	*mock.Call
}

// GetCreateInput is a helper method to define mock.On call
func (_e *Handler_Expecter) GetCreateInput() *Handler_GetCreateInput_Call {
	return &Handler_GetCreateInput_Call{Call: _e.mock.On("GetCreateInput")}
}

func (_c *Handler_GetCreateInput_Call) Run(run func()) *Handler_GetCreateInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_GetCreateInput_Call) Return(_a0 interface{}) *Handler_GetCreateInput_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_GetCreateInput_Call) RunAndReturn(run func() interface{}) *Handler_GetCreateInput_Call {
	_c.Call.Return(run)
	return _c
}

// GetModel provides a mock function with given fields:
func (_m *Handler) GetModel() db_repo.ModelBased {
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

// Handler_GetModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModel'
type Handler_GetModel_Call struct {
	*mock.Call
}

// GetModel is a helper method to define mock.On call
func (_e *Handler_Expecter) GetModel() *Handler_GetModel_Call {
	return &Handler_GetModel_Call{Call: _e.mock.On("GetModel")}
}

func (_c *Handler_GetModel_Call) Run(run func()) *Handler_GetModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_GetModel_Call) Return(_a0 db_repo.ModelBased) *Handler_GetModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_GetModel_Call) RunAndReturn(run func() db_repo.ModelBased) *Handler_GetModel_Call {
	_c.Call.Return(run)
	return _c
}

// GetRepository provides a mock function with given fields:
func (_m *Handler) GetRepository() crud.Repository {
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

// Handler_GetRepository_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRepository'
type Handler_GetRepository_Call struct {
	*mock.Call
}

// GetRepository is a helper method to define mock.On call
func (_e *Handler_Expecter) GetRepository() *Handler_GetRepository_Call {
	return &Handler_GetRepository_Call{Call: _e.mock.On("GetRepository")}
}

func (_c *Handler_GetRepository_Call) Run(run func()) *Handler_GetRepository_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_GetRepository_Call) Return(_a0 crud.Repository) *Handler_GetRepository_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_GetRepository_Call) RunAndReturn(run func() crud.Repository) *Handler_GetRepository_Call {
	_c.Call.Return(run)
	return _c
}

// GetUpdateInput provides a mock function with given fields:
func (_m *Handler) GetUpdateInput() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Handler_GetUpdateInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUpdateInput'
type Handler_GetUpdateInput_Call struct {
	*mock.Call
}

// GetUpdateInput is a helper method to define mock.On call
func (_e *Handler_Expecter) GetUpdateInput() *Handler_GetUpdateInput_Call {
	return &Handler_GetUpdateInput_Call{Call: _e.mock.On("GetUpdateInput")}
}

func (_c *Handler_GetUpdateInput_Call) Run(run func()) *Handler_GetUpdateInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_GetUpdateInput_Call) Return(_a0 interface{}) *Handler_GetUpdateInput_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_GetUpdateInput_Call) RunAndReturn(run func() interface{}) *Handler_GetUpdateInput_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, qb, apiView
func (_m *Handler) List(ctx context.Context, qb *db_repo.QueryBuilder, apiView string) (interface{}, error) {
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

// Handler_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Handler_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - apiView string
func (_e *Handler_Expecter) List(ctx interface{}, qb interface{}, apiView interface{}) *Handler_List_Call {
	return &Handler_List_Call{Call: _e.mock.On("List", ctx, qb, apiView)}
}

func (_c *Handler_List_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, apiView string)) *Handler_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(string))
	})
	return _c
}

func (_c *Handler_List_Call) Return(out interface{}, err error) *Handler_List_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *Handler_List_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, string) (interface{}, error)) *Handler_List_Call {
	_c.Call.Return(run)
	return _c
}

// TransformCreate provides a mock function with given fields: ctx, input, model
func (_m *Handler) TransformCreate(ctx context.Context, input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(ctx, input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_TransformCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformCreate'
type Handler_TransformCreate_Call struct {
	*mock.Call
}

// TransformCreate is a helper method to define mock.On call
//   - ctx context.Context
//   - input interface{}
//   - model db_repo.ModelBased
func (_e *Handler_Expecter) TransformCreate(ctx interface{}, input interface{}, model interface{}) *Handler_TransformCreate_Call {
	return &Handler_TransformCreate_Call{Call: _e.mock.On("TransformCreate", ctx, input, model)}
}

func (_c *Handler_TransformCreate_Call) Run(run func(ctx context.Context, input interface{}, model db_repo.ModelBased)) *Handler_TransformCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Handler_TransformCreate_Call) Return(err error) *Handler_TransformCreate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Handler_TransformCreate_Call) RunAndReturn(run func(context.Context, interface{}, db_repo.ModelBased) error) *Handler_TransformCreate_Call {
	_c.Call.Return(run)
	return _c
}

// TransformOutput provides a mock function with given fields: ctx, model, apiView
func (_m *Handler) TransformOutput(ctx context.Context, model db_repo.ModelBased, apiView string) (interface{}, error) {
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

// Handler_TransformOutput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformOutput'
type Handler_TransformOutput_Call struct {
	*mock.Call
}

// TransformOutput is a helper method to define mock.On call
//   - ctx context.Context
//   - model db_repo.ModelBased
//   - apiView string
func (_e *Handler_Expecter) TransformOutput(ctx interface{}, model interface{}, apiView interface{}) *Handler_TransformOutput_Call {
	return &Handler_TransformOutput_Call{Call: _e.mock.On("TransformOutput", ctx, model, apiView)}
}

func (_c *Handler_TransformOutput_Call) Run(run func(ctx context.Context, model db_repo.ModelBased, apiView string)) *Handler_TransformOutput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_repo.ModelBased), args[2].(string))
	})
	return _c
}

func (_c *Handler_TransformOutput_Call) Return(output interface{}, err error) *Handler_TransformOutput_Call {
	_c.Call.Return(output, err)
	return _c
}

func (_c *Handler_TransformOutput_Call) RunAndReturn(run func(context.Context, db_repo.ModelBased, string) (interface{}, error)) *Handler_TransformOutput_Call {
	_c.Call.Return(run)
	return _c
}

// TransformUpdate provides a mock function with given fields: ctx, input, model
func (_m *Handler) TransformUpdate(ctx context.Context, input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(ctx, input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_TransformUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransformUpdate'
type Handler_TransformUpdate_Call struct {
	*mock.Call
}

// TransformUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - input interface{}
//   - model db_repo.ModelBased
func (_e *Handler_Expecter) TransformUpdate(ctx interface{}, input interface{}, model interface{}) *Handler_TransformUpdate_Call {
	return &Handler_TransformUpdate_Call{Call: _e.mock.On("TransformUpdate", ctx, input, model)}
}

func (_c *Handler_TransformUpdate_Call) Run(run func(ctx context.Context, input interface{}, model db_repo.ModelBased)) *Handler_TransformUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Handler_TransformUpdate_Call) Return(err error) *Handler_TransformUpdate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Handler_TransformUpdate_Call) RunAndReturn(run func(context.Context, interface{}, db_repo.ModelBased) error) *Handler_TransformUpdate_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
