package service

import (
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StudentServiceSuite struct {
	suite.Suite
	msr     *mocks.StudentRepository
	service *Student
}

func (suite *StudentServiceSuite) SetupTest() {
	suite.msr = &mocks.StudentRepository{}
	suite.service = &Student{suite.msr}
}

func (suite *StudentServiceSuite) TestGetStudent() {
	s := factory.NewStudent()
	suite.msr.On("GetStudent", 1).Return(s, nil)
	rs, err := suite.service.Get(1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), s, rs)
}

func (suite *StudentServiceSuite) TestDeleteStudent() {
	suite.msr.On("DeleteStudent", 1).Return(nil)
	err := suite.service.Delete(1)
	suite.Nil(err)
	suite.msr.AssertExpectations(suite.T())
}

func (suite *StudentServiceSuite) TestEditStudent() {
	s := factory.NewStudent()

	suite.msr.On("EditStudent", &s).Return(nil)

	err := suite.service.Edit(&s)
	suite.Nil(err)
	suite.msr.AssertExpectations(suite.T())
}

func TestStudentService(t *testing.T) {
	suite.Run(t, new(StudentServiceSuite))
}

func (suite *StudentServiceSuite) TestAddStudent() {
	s := factory.NewStudent()

	suite.msr.On("AddStudent", &s).Return(nil)

	err := suite.service.Add(&s)
	suite.Nil(err)
	suite.msr.AssertExpectations(suite.T())
}

func (suite *StudentServiceSuite) TestAddStudentExceedingCreditLimit() {
	s := factory.NewStudent()
	s.Courses = []demo.Course{
		{
			Credit: 40,
		},
	}

	err := suite.service.Add(&s)
	suite.Equal(err, OverCreditErr, "Should return OverCreditErr")
	suite.msr.AssertNotCalled(suite.T(), "AddStudent")

}
