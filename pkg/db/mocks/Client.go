// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/justtrackio/gosoline/pkg/db"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *Client) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type Client_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Exec(ctx interface{}, query interface{}, args ...interface{}) *Client_Exec_Call {
	return &Client_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_Exec_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Exec_Call) Return(_a0 sql.Result, _a1 error) *Client_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *Client_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, dest, query, args
func (_m *Client) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Client_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Get(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *Client_Get_Call {
	return &Client_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *Client_Get_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *Client_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Get_Call) Return(_a0 error) *Client_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Get_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *Client_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetResult provides a mock function with given fields: ctx, query, args
func (_m *Client) GetResult(ctx context.Context, query string, args ...interface{}) (*db.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *db.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*db.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *db.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResult'
type Client_GetResult_Call struct {
	*mock.Call
}

// GetResult is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) GetResult(ctx interface{}, query interface{}, args ...interface{}) *Client_GetResult_Call {
	return &Client_GetResult_Call{Call: _e.mock.On("GetResult",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_GetResult_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_GetResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_GetResult_Call) Return(_a0 *db.Result, _a1 error) *Client_GetResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetResult_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*db.Result, error)) *Client_GetResult_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleScalarValue provides a mock function with given fields: ctx, query, args
func (_m *Client) GetSingleScalarValue(ctx context.Context, query string, args ...interface{}) (int, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (int, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) int); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetSingleScalarValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleScalarValue'
type Client_GetSingleScalarValue_Call struct {
	*mock.Call
}

// GetSingleScalarValue is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) GetSingleScalarValue(ctx interface{}, query interface{}, args ...interface{}) *Client_GetSingleScalarValue_Call {
	return &Client_GetSingleScalarValue_Call{Call: _e.mock.On("GetSingleScalarValue",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_GetSingleScalarValue_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_GetSingleScalarValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_GetSingleScalarValue_Call) Return(_a0 int, _a1 error) *Client_GetSingleScalarValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetSingleScalarValue_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (int, error)) *Client_GetSingleScalarValue_Call {
	_c.Call.Return(run)
	return _c
}

// Prepare provides a mock function with given fields: ctx, query
func (_m *Client) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	ret := _m.Called(ctx, query)

	var r0 *sql.Stmt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Stmt, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Stmt); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Prepare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prepare'
type Client_Prepare_Call struct {
	*mock.Call
}

// Prepare is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
func (_e *Client_Expecter) Prepare(ctx interface{}, query interface{}) *Client_Prepare_Call {
	return &Client_Prepare_Call{Call: _e.mock.On("Prepare", ctx, query)}
}

func (_c *Client_Prepare_Call) Run(run func(ctx context.Context, query string)) *Client_Prepare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_Prepare_Call) Return(_a0 *sql.Stmt, _a1 error) *Client_Prepare_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Prepare_Call) RunAndReturn(run func(context.Context, string) (*sql.Stmt, error)) *Client_Prepare_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, query, args
func (_m *Client) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type Client_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Query(ctx interface{}, query interface{}, args ...interface{}) *Client_Query_Call {
	return &Client_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_Query_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Query_Call) Return(_a0 *sql.Rows, _a1 error) *Client_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *Client_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, query, args
func (_m *Client) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Client_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type Client_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) QueryRow(ctx interface{}, query interface{}, args ...interface{}) *Client_QueryRow_Call {
	return &Client_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_QueryRow_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_QueryRow_Call) Return(_a0 *sql.Row) *Client_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sql.Row) *Client_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// Queryx provides a mock function with given fields: ctx, query, args
func (_m *Client) Queryx(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sqlx.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Queryx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Queryx'
type Client_Queryx_Call struct {
	*mock.Call
}

// Queryx is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Queryx(ctx interface{}, query interface{}, args ...interface{}) *Client_Queryx_Call {
	return &Client_Queryx_Call{Call: _e.mock.On("Queryx",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Client_Queryx_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Client_Queryx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Queryx_Call) Return(_a0 *sqlx.Rows, _a1 error) *Client_Queryx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Queryx_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sqlx.Rows, error)) *Client_Queryx_Call {
	_c.Call.Return(run)
	return _c
}

// Select provides a mock function with given fields: ctx, dest, query, args
func (_m *Client) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type Client_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *Client_Expecter) Select(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *Client_Select_Call {
	return &Client_Select_Call{Call: _e.mock.On("Select",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *Client_Select_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *Client_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Select_Call) Return(_a0 error) *Client_Select_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Select_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *Client_Select_Call {
	_c.Call.Return(run)
	return _c
}

// WithTx provides a mock function with given fields: ctx, ops, do
func (_m *Client) WithTx(ctx context.Context, ops *sql.TxOptions, do func(context.Context, *sql.Tx) error) error {
	ret := _m.Called(ctx, ops, do)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions, func(context.Context, *sql.Tx) error) error); ok {
		r0 = rf(ctx, ops, do)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type Client_WithTx_Call struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//   - ctx context.Context
//   - ops *sql.TxOptions
//   - do func(context.Context , *sql.Tx) error
func (_e *Client_Expecter) WithTx(ctx interface{}, ops interface{}, do interface{}) *Client_WithTx_Call {
	return &Client_WithTx_Call{Call: _e.mock.On("WithTx", ctx, ops, do)}
}

func (_c *Client_WithTx_Call) Run(run func(ctx context.Context, ops *sql.TxOptions, do func(context.Context, *sql.Tx) error)) *Client_WithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sql.TxOptions), args[2].(func(context.Context, *sql.Tx) error))
	})
	return _c
}

func (_c *Client_WithTx_Call) Return(_a0 error) *Client_WithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_WithTx_Call) RunAndReturn(run func(context.Context, *sql.TxOptions, func(context.Context, *sql.Tx) error) error) *Client_WithTx_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
