package mocks

import (
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
	"github.com/stretchr/testify/mock"
)

type DatabaseHelper struct {
	mock.Mock
}
// Client provides a mock function with given fields:
func (_m *DatabaseHelper) Client() mongos.Client {
	ret := _m.Called()

	var r0 mongos.Client
	if rf, ok := ret.Get(0).(func() mongos.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongos.Client)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: name
func (_m *DatabaseHelper) Collection(name string) mongos.Collection {
	ret := _m.Called(name)

	var r0 mongos.Collection
	if rf, ok := ret.Get(0).(func(string) mongos.Collection); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongos.Collection)
		}
	}

	return r0
}
