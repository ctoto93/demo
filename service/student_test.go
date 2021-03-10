package service_test

import (
	"testing"

	"github.com/ctoto93/demo/service"
	"github.com/ctoto93/demo/test/factory"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentServiceSuite struct {
	suite.Suite
	service          *service.Student
	repo             service.Repository
	cleanUpDb        func()
	disconnectClient func()
}

func (suite *StudentServiceSuite) TearDownTest() {
	suite.cleanUpDb()
}

func (suite *StudentServiceSuite) TearDownSuite() {
	suite.disconnectClient()
}

func TestStudentMongoService(t *testing.T) {
	client, db, repo := InitTestMongoRepo(t)
	serv := service.NewStudent(repo)
	s := StudentServiceSuite{
		service:          serv,
		repo:             repo,
		cleanUpDb:        buildMongoCleanUpFunc(db),
		disconnectClient: buildMongoClientDisconnectFunc(client),
	}
	suite.Run(t, &s)
}

func TestStudentSQLiteService(t *testing.T) {
	tx, repo := InitTestSQLiteRepo(t)
	serv := service.NewStudent(repo)
	s := StudentServiceSuite{
		service:          serv,
		repo:             repo,
		cleanUpDb:        buildSQLCleanUpFunc(tx),
		disconnectClient: func() {},
	}
	suite.Run(t, &s)
}

func (suite *StudentServiceSuite) TestGetStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.Get(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *StudentServiceSuite) TestAddStudent() {

	expected := factory.NewStudentWithCourses(1)
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&expected)
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

func (suite *StudentServiceSuite) TestAddStudentExceedingCreditLimit() {
	expected := factory.NewStudentWithCourses(1)
	expected.Courses[0].Credit = 40
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}

	err := suite.service.Add(&expected)
	suite.Require().Equal(service.OverCreditErr, err, "Should return OverCreditErr")

}

func (suite *StudentServiceSuite) TestEditStudent() {

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

	err = suite.service.Edit(&expected)
	suite.Require().Nil(err)

	actual, err := suite.repo.GetStudent(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

	newCourse, err = suite.repo.GetCourse(newCourse.Id)
	suite.Require().Nil(err)
	suite.Require().True(newCourse.HasStudent(expected), "Added courses should have the the student")

	expected.Courses = expected.Courses[:len(expected.Courses)-1]
	err = suite.service.Edit(&expected)
	suite.Require().Nil(err)

	removedCourse, err := suite.repo.GetCourse(newCourse.Id)

	suite.Require().Nil(err)
	suite.Require().False(removedCourse.HasStudent(expected), "Removed courses should not have the student")

}

func (suite *StudentServiceSuite) TestDeleteStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	err = suite.service.Delete(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetStudent(expected.Id)
	suite.Require().Equal(err, mongo.ErrNoDocuments)

}
