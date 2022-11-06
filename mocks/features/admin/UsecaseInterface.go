// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	admin "be12/mentutor/features/admin"

	mock "github.com/stretchr/testify/mock"
)

// UsecaseInterface is an autogenerated mock type for the UsecaseInterface type
type UsecaseInterface struct {
	mock.Mock
}

// AddNewClass provides a mock function with given fields: input, role
func (_m *UsecaseInterface) AddNewClass(input admin.ClassCore, role string) error {
	ret := _m.Called(input, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.ClassCore, string) error); ok {
		r0 = rf(input, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUser provides a mock function with given fields: input, role
func (_m *UsecaseInterface) AddUser(input admin.UserCore, role string) (admin.UserCore, error) {
	ret := _m.Called(input, role)

	var r0 admin.UserCore
	if rf, ok := ret.Get(0).(func(admin.UserCore, string) admin.UserCore); ok {
		r0 = rf(input, role)
	} else {
		r0 = ret.Get(0).(admin.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.UserCore, string) error); ok {
		r1 = rf(input, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteClass provides a mock function with given fields: id, role
func (_m *UsecaseInterface) DeleteClass(id uint, role string) error {
	ret := _m.Called(id, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(id, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id, role
func (_m *UsecaseInterface) DeleteUser(id uint, role string) error {
	ret := _m.Called(id, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(id, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllClass provides a mock function with given fields: role
func (_m *UsecaseInterface) GetAllClass(role string) ([]admin.ClassCore, error) {
	ret := _m.Called(role)

	var r0 []admin.ClassCore
	if rf, ok := ret.Get(0).(func(string) []admin.ClassCore); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.ClassCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUser provides a mock function with given fields: role
func (_m *UsecaseInterface) GetAllUser(role string) ([]admin.UserCore, []admin.UserCore, error) {
	ret := _m.Called(role)

	var r0 []admin.UserCore
	if rf, ok := ret.Get(0).(func(string) []admin.UserCore); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.UserCore)
		}
	}

	var r1 []admin.UserCore
	if rf, ok := ret.Get(1).(func(string) []admin.UserCore); ok {
		r1 = rf(role)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]admin.UserCore)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(role)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSingleUser provides a mock function with given fields: id, role
func (_m *UsecaseInterface) GetSingleUser(id uint, role string) (admin.UserCore, error) {
	ret := _m.Called(id, role)

	var r0 admin.UserCore
	if rf, ok := ret.Get(0).(func(uint, string) admin.UserCore); ok {
		r0 = rf(id, role)
	} else {
		r0 = ret.Get(0).(admin.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(id, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClass provides a mock function with given fields: input, role
func (_m *UsecaseInterface) UpdateClass(input admin.ClassCore, role string) (admin.ClassCore, error) {
	ret := _m.Called(input, role)

	var r0 admin.ClassCore
	if rf, ok := ret.Get(0).(func(admin.ClassCore, string) admin.ClassCore); ok {
		r0 = rf(input, role)
	} else {
		r0 = ret.Get(0).(admin.ClassCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.ClassCore, string) error); ok {
		r1 = rf(input, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserAdmin provides a mock function with given fields: input, role
func (_m *UsecaseInterface) UpdateUserAdmin(input admin.UserCore, role string) (admin.UserCore, error) {
	ret := _m.Called(input, role)

	var r0 admin.UserCore
	if rf, ok := ret.Get(0).(func(admin.UserCore, string) admin.UserCore); ok {
		r0 = rf(input, role)
	} else {
		r0 = ret.Get(0).(admin.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.UserCore, string) error); ok {
		r1 = rf(input, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecaseInterface creates a new instance of UsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecaseInterface(t mockConstructorTestingTNewUsecaseInterface) *UsecaseInterface {
	mock := &UsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
