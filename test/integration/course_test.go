package integration_test

import (
	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
)

func (suite *ServiceSuite) TestGetCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.GetCourse(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *ServiceSuite) TestAddCourse() {

	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.AddCourse(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetCourse(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	for i := range expected.Students {
		s, err := suite.repo.GetStudent(expected.Students[i].Id)
		suite.Require().Nil(err)
		suite.Require().True(s.HasCourse(expected), "Should add course to the respective student")
	}

}

func (suite *ServiceSuite) TestAddLessThanMinStudents() {
	c := factory.NewCourseWithStudents(demo.MinNumOfStudents - 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.AddCourse(&c)
	suite.Equal(demo.InsufficientStudentsErr, err, "Should return InsufficientStudents Err")

}

func (suite *ServiceSuite) TestAddMoreThanMaxStudents() {

	c := factory.NewCourseWithStudents(demo.MaxNumOfStudents + 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.AddCourse(&c)
	suite.Require().Equal(demo.ExceedingStudentsErr, err, "Should return ExceedingStudents Err")

}

func (suite *ServiceSuite) TestEditCourse() {

	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	expected.Name = "Edit"
	expected.Credit -= 1

	newStudent := factory.NewStudent()
	err = suite.repo.AddStudent(&newStudent)
	suite.Require().Nil(err)
	expected.Students = append(expected.Students, newStudent)

	err = suite.service.EditCourse(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetCourse(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	newStudent, err = suite.repo.GetStudent(newStudent.Id)
	suite.Require().Nil(err)
	suite.Require().True(newStudent.HasCourse(expected), "Added students should have the course")

	expected.Students = expected.Students[:len(expected.Students)-1]
	err = suite.service.EditCourse(&expected)
	suite.Require().Nil(err)

	removedStudent, err := suite.repo.GetStudent(newStudent.Id)

	suite.Require().Nil(err)
	suite.Require().False(removedStudent.HasCourse(expected), "Removed students should not have the course")

}

func (suite *ServiceSuite) TestEditCourseLessThanMinStudents() {

	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	expected.Students = expected.Students[:demo.MinNumOfStudents-1]

	err = suite.service.EditCourse(&expected)
	suite.Equal(demo.InsufficientStudentsErr, err, "Should return service.InsufficientStudentsErr")

}

func (suite *ServiceSuite) TestEditCourseMoreThanMaxStudents() {

	expected := factory.NewCourseWithStudents(demo.MaxNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}

	newStud := factory.NewStudent()
	err := suite.repo.AddStudent(&newStud)
	suite.Require().Nil(err)

	expected.Students = append(expected.Students, newStud)
	err = suite.service.EditCourse(&expected)
	suite.Equal(demo.ExceedingStudentsErr, err, "Should return service.ExceedingStudentsErr")

}

func (suite *ServiceSuite) TestDeleteCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	err = suite.service.DeleteCourse(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetCourse(expected.Id)
	suite.Require().NotEmpty(err)

}
