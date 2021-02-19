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
	s := factory.NewCourse()

	suite.repo.On("AddCourse", &s).Return(nil)

	err := suite.service.Add(&s)
	suite.Nil(err)
	suite.repo.AssertExpectations(suite.T())
}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseServiceSuite))
}
