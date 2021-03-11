package integration_test

import (
	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
)

func (suite *ServiceSuite) TestGetStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.GetStudent(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *ServiceSuite) TestAddStudent() {

	expected := factory.NewStudentWithCourses(1)
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}
	err := suite.service.AddStudent(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetStudent(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	for i := range expected.Courses {
		c, err := suite.repo.GetCourse(expected.Courses[i].Id)
		suite.Require().Nil(err)
		suite.Require().True(c.HasStudent(expected), "Should add student to the respective course")
	}

}

func (suite *ServiceSuite) TestAddStudentExceedingCreditLimit() {
	expected := factory.NewStudentWithCourses(1)
	expected.Courses[0].Credit = 40
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}

	err := suite.service.AddStudent(&expected)
	suite.Require().Equal(demo.OverCreditErr, err, "Should return OverCreditErr")

}

func (suite *ServiceSuite) TestEditStudent() {

	expected := factory.NewStudentWithCourses(1)
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	expected.Name = "Edit"
	expected.Age -= 1

	newCourse := factory.NewCourse()
	err = suite.repo.AddCourse(&newCourse)
	suite.Require().Nil(err)
	expected.Courses = append(expected.Courses, newCourse)

	err = suite.service.EditStudent(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetStudent(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	newCourse, err = suite.repo.GetCourse(newCourse.Id)
	suite.Require().Nil(err)
	suite.Require().True(newCourse.HasStudent(expected), "Added courses should have the the student")

	expected.Courses = expected.Courses[:len(expected.Courses)-1]
	err = suite.service.EditStudent(&expected)
	suite.Require().Nil(err)

	removedCourse, err := suite.repo.GetCourse(newCourse.Id)

	suite.Require().Nil(err)
	suite.Require().False(removedCourse.HasStudent(expected), "Removed courses should not have the student")

}

func (suite *ServiceSuite) TestDeleteStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	err = suite.service.DeleteStudent(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetStudent(expected.Id)
	suite.Require().NotEmpty(err)

}
