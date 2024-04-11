// Code generated by mockery v2.42.0. DO NOT EDIT.

package interfaces

import mock "github.com/stretchr/testify/mock"

// MockHasher is an autogenerated mock type for the Hasher type
type MockHasher struct {
	mock.Mock
}

type MockHasher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHasher) EXPECT() *MockHasher_Expecter {
	return &MockHasher_Expecter{mock: &_m.Mock}
}

// Compare provides a mock function with given fields: key, toCompare
func (_m *MockHasher) Compare(key string, toCompare string) bool {
	ret := _m.Called(key, toCompare)

	if len(ret) == 0 {
		panic("no return value specified for Compare")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(key, toCompare)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockHasher_Compare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compare'
type MockHasher_Compare_Call struct {
	*mock.Call
}

// Compare is a helper method to define mock.On call
//   - key string
//   - toCompare string
func (_e *MockHasher_Expecter) Compare(key interface{}, toCompare interface{}) *MockHasher_Compare_Call {
	return &MockHasher_Compare_Call{Call: _e.mock.On("Compare", key, toCompare)}
}

func (_c *MockHasher_Compare_Call) Run(run func(key string, toCompare string)) *MockHasher_Compare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockHasher_Compare_Call) Return(_a0 bool) *MockHasher_Compare_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHasher_Compare_Call) RunAndReturn(run func(string, string) bool) *MockHasher_Compare_Call {
	_c.Call.Return(run)
	return _c
}

// Hash provides a mock function with given fields: key
func (_m *MockHasher) Hash(key string) (string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Hash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHasher_Hash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hash'
type MockHasher_Hash_Call struct {
	*mock.Call
}

// Hash is a helper method to define mock.On call
//   - key string
func (_e *MockHasher_Expecter) Hash(key interface{}) *MockHasher_Hash_Call {
	return &MockHasher_Hash_Call{Call: _e.mock.On("Hash", key)}
}

func (_c *MockHasher_Hash_Call) Run(run func(key string)) *MockHasher_Hash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHasher_Hash_Call) Return(_a0 string, _a1 error) *MockHasher_Hash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHasher_Hash_Call) RunAndReturn(run func(string) (string, error)) *MockHasher_Hash_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHasher creates a new instance of MockHasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHasher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHasher {
	mock := &MockHasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
