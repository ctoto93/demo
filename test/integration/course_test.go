package integration_test

import (
	"testing"

	"github.com/ctoto93/demo/service"
	"github.com/ctoto93/demo/test/factory"
	"github.com/stretchr/testify/suite"
)

type CourseServiceSuite struct {
	suite.Suite
	repo             service.Repository
	service          *service.Course
	cleanUpDb        func()
	disconnectClient func()
}

func (suite *CourseServiceSuite) TearDownTest() {
	suite.cleanUpDb()
}

func (suite *CourseServiceSuite) TearDownSuite() {
	suite.disconnectClient()
}

func TestCourseMongoService(t *testing.T) {
	client, db, repo := InitTestMongoRepo(t)
	serv := service.NewCourse(repo)
	s := CourseServiceSuite{
		service:          serv,
		repo:             repo,
		cleanUpDb:        buildMongoCleanUpFunc(db),
		disconnectClient: buildMongoClientDisconnectFunc(client),
	}
	suite.Run(t, &s)
}

func TestCourseSQLiteService(t *testing.T) {
	tx, repo := InitTestSQLiteRepo(t)
	serv := service.NewCourse(repo)
	s := CourseServiceSuite{
		service:          serv,
		repo:             repo,
		cleanUpDb:        buildSQLCleanUpFunc(tx),
		disconnectClient: buildSQLClientDisconnect(),
	}
	suite.Run(t, &s)
}

func (suite *CourseServiceSuite) TestGetCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.Get(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *CourseServiceSuite) TestAddCourse() {

	expected := factory.NewCourseWithStudents(service.MinNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&expected)
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

func (suite *CourseServiceSuite) TestAddLessThanMinStudents() {
	c := factory.NewCourseWithStudents(service.MinNumOfStudents - 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&c)
	suite.Equal(service.InsufficientStudentsErr, err, "Should return InsufficientStudents Err")

}

func (suite *CourseServiceSuite) TestAddMoreThanMaxStudents() {

	c := factory.NewCourseWithStudents(service.MaxNumOfStudents + 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&c)
	suite.Require().Equal(service.ExceedingStudentsErr, err, "Should return ExceedingStudents Err")

}

func (suite *CourseServiceSuite) TestEditCourse() {

	expected := factory.NewCourseWithStudents(service.MinNumOfStudents)
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

	err = suite.service.Edit(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetCourse(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	newStudent, err = suite.repo.GetStudent(newStudent.Id)
	suite.Require().Nil(err)
	suite.Require().True(newStudent.HasCourse(expected), "Added students should have the course")

	expected.Students = expected.Students[:len(expected.Students)-1]
	err = suite.service.Edit(&expected)
	suite.Require().Nil(err)

	removedStudent, err := suite.repo.GetStudent(newStudent.Id)

	suite.Require().Nil(err)
	suite.Require().False(removedStudent.HasCourse(expected), "Removed students should not have the course")

}

func (suite *CourseServiceSuite) TestEditCourseLessThanMinStudents() {

	expected := factory.NewCourseWithStudents(service.MinNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	expected.Students = expected.Students[:service.MinNumOfStudents-1]

	err = suite.service.Edit(&expected)
	suite.Equal(service.InsufficientStudentsErr, err, "Should return service.InsufficientStudentsErr")

}

func (suite *CourseServiceSuite) TestEditCourseMoreThanMaxStudents() {

	expected := factory.NewCourseWithStudents(service.MaxNumOfStudents)
	for i := range expected.Students {
		err := suite.repo.AddStudent(&expected.Students[i])
		suite.Require().Nil(err)
	}

	newStud := factory.NewStudent()
	err := suite.repo.AddStudent(&newStud)
	suite.Require().Nil(err)

	expected.Students = append(expected.Students, newStud)
	err = suite.service.Edit(&expected)
	suite.Equal(service.ExceedingStudentsErr, err, "Should return service.ExceedingStudentsErr")

}

func (suite *CourseServiceSuite) TestDeleteCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	err = suite.service.Delete(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetCourse(expected.Id)
	suite.Require().NotEmpty(err)

}
