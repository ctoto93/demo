package service

import (
	"testing"

	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/stretchr/testify/suite"
)

type CourseServiceSuite struct {
	suite.Suite
	repo    *mocks.CourseRepository
	service *Course
}

func (suite *CourseServiceSuite) SetupTest() {
	suite.repo = &mocks.CourseRepository{}
	suite.service = &Course{suite.repo}
}

func (suite *CourseServiceSuite) TestGetCourse() {
	c := factory.NewCourse()
	suite.repo.On("GetCourse", 1).Return(c, nil)
	rc, err := suite.service.Get(1)
	suite.Nil(err)
	suite.Equal(c, rc)
}

func (suite *CourseServiceSuite) TestDeleteCourse() {
	suite.repo.On("DeleteCourse", 1).Return(nil)
	err := suite.service.Delete(1)
	suite.Nil(err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *CourseServiceSuite) TestEditCourse() {
	c := factory.NewCourse()

	suite.repo.On("EditCourse", &c).Return(nil)

	err := suite.service.Edit(&c)
	suite.Nil(err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *CourseServiceSuite) TestAddCourse() {
	c := factory.NewCourse()
	c.Students = factory.NewStudents(5)

	suite.repo.On("AddCourse", &c).Return(nil)

	err := suite.service.Add(&c)
	suite.Nil(err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *CourseServiceSuite) TestAddCourseLessThanFiveStudents() {
	c := factory.NewCourse()
	c.Students = factory.NewStudents(3)

	suite.repo.On("AddCourse", &c).Return(nil)

	err := suite.service.Add(&c)
	suite.Equal(InsufficientStudentsErr, err, "Should return InsufficientStudents Err")
	suite.repo.AssertNotCalled(suite.T(), "AddStudent")

}

func (suite *CourseServiceSuite) TestAddCourseMoreThanThirtyStudents() {
	c := factory.NewCourse()
	c.Students = factory.NewStudents(31)

	suite.repo.On("AddCourse", &c).Return(nil)

	err := suite.service.Add(&c)
	suite.Equal(ExceedingStudentsErr, err, "Should return ExceedingStudentsErr Err")
	suite.repo.AssertNotCalled(suite.T(), "AddStudent")

}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseServiceSuite))
}
