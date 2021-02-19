// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	demo "github.com/ctoto93/demo"
	mock "github.com/stretchr/testify/mock"
)

// StudentRepository is an autogenerated mock type for the StudentRepository type
type StudentRepository struct {
	mock.Mock
}

// AddStudent provides a mock function with given fields: s
func (_m *StudentRepository) AddStudent(s *demo.Student) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Student) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStudent provides a mock function with given fields: id
func (_m *StudentRepository) DeleteStudent(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditStudent provides a mock function with given fields: s
func (_m *StudentRepository) EditStudent(s *demo.Student) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*demo.Student) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStudent provides a mock function with given fields: id
func (_m *StudentRepository) GetStudent(id int) (demo.Student, error) {
	ret := _m.Called(id)

	var r0 demo.Student
	if rf, ok := ret.Get(0).(func(int) demo.Student); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(demo.Student)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}