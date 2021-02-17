package service

import (
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	assert := assert.New(t)
	msr := &mocks.StudentRepository{}
	serv := &Student{msr}

	s := demo.Student{
		Name: "dummy",
		Age:  20,
	}

	msr.On("GetStudent", 1).Return(s, nil)

	rs, err := serv.Get(1)
	assert.Equal(s, rs)
	assert.Nil(err)
	msr.AssertExpectations(t)
}

func TestDeleteStudent(t *testing.T) {
	assert := assert.New(t)
	msr := &mocks.StudentRepository{}
	serv := &Student{msr}

	msr.On("DeleteStudent", 1).Return(nil)

	err := serv.Delete(1)
	assert.Nil(err)
	msr.AssertExpectations(t)
}

func TestEditStudent(t *testing.T) {
	assert := assert.New(t)
	msr := &mocks.StudentRepository{}
	serv := &Student{msr}

	s := &demo.Student{
		Name: "dummy",
		Age:  20,
	}

	msr.On("EditStudent", s).Return(nil)

	err := serv.Edit(s)
	assert.Nil(err)
	msr.AssertExpectations(t)
}

func TestAddStudent(t *testing.T) {
	assert := assert.New(t)
	msr := &mocks.StudentRepository{}
	serv := &Student{msr}

	s := &demo.Student{
		Name: "dummy",
		Age:  20,
	}

	msr.On("AddStudent", s).Return(nil)

	err := serv.Add(s)
	assert.Nil(err)
	msr.AssertExpectations(t)
}

func TestAddStudentExceedingCreditLimit(t *testing.T) {
	assert := assert.New(t)
	msr := &mocks.StudentRepository{}
	serv := &Student{msr}
	s := demo.Student{
		Name: "dummy",
		Age:  20,
		Courses: []demo.Course{
			{
				Credit: 40,
			},
		},
	}

	err := serv.Add(&s)
	assert.Equal(err, OverCreditErr, "Should return OverCreditErr")
	msr.AssertNotCalled(t, "AddStudent")

}
