// Code generated by mockery. DO NOT EDIT.

package schema

import mock "github.com/stretchr/testify/mock"

// Blueprint is an autogenerated mock type for the Blueprint type
type Blueprint struct {
	mock.Mock
}

type Blueprint_Expecter struct {
	mock *mock.Mock
}

func (_m *Blueprint) EXPECT() *Blueprint_Expecter {
	return &Blueprint_Expecter{mock: &_m.Mock}
}

// NewBlueprint creates a new instance of Blueprint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlueprint(t interface {
	mock.TestingT
	Cleanup(func())
}) *Blueprint {
	mock := &Blueprint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
