// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	demo "github.com/ctoto93/demo"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddCourse provides a mock function with given fields: c
func (_m *Repository) AddCourse(c *demo.Course) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Course) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddStudent provides a mock function with given fields: s
func (_m *Repository) AddStudent(s *demo.Student) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Student) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCourse provides a mock function with given fields: id
func (_m *Repository) DeleteCourse(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStudent provides a mock function with given fields: id
func (_m *Repository) DeleteStudent(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditCourse provides a mock function with given fields: c
func (_m *Repository) EditCourse(c *demo.Course) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Course) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditStudent provides a mock function with given fields: s
func (_m *Repository) EditStudent(s *demo.Student) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Student) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCourse provides a mock function with given fields: id
func (_m *Repository) GetCourse(id string) (demo.Course, error) {
	ret := _m.Called(id)

	var r0 demo.Course
	if rf, ok := ret.Get(0).(func(string) demo.Course); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(demo.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStudent provides a mock function with given fields: id
func (_m *Repository) GetStudent(id string) (demo.Student, error) {
	ret := _m.Called(id)

	var r0 demo.Student
	if rf, ok := ret.Get(0).(func(string) demo.Student); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(demo.Student)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}