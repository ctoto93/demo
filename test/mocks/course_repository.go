// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	demo "github.com/ctoto93/demo"
	mock "github.com/stretchr/testify/mock"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// AddCourse provides a mock function with given fields: s
func (_m *CourseRepository) AddCourse(s *demo.Course) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Course) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCourse provides a mock function with given fields: id
func (_m *CourseRepository) DeleteCourse(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditCourse provides a mock function with given fields: s
func (_m *CourseRepository) EditCourse(s *demo.Course) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Course) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCourse provides a mock function with given fields: id
func (_m *CourseRepository) GetCourse(id int) (demo.Course, error) {
	ret := _m.Called(id)

	var r0 demo.Course
	if rf, ok := ret.Get(0).(func(int) demo.Course); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(demo.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}