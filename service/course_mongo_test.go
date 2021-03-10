package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/ctoto93/demo/service"
	"github.com/ctoto93/demo/test/factory"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseMongoServiceSuite struct {
	suite.Suite
	client  *mongo.Client
	db      *mongo.Database
	repo    service.Repository
	service *service.Course
}

func (suite *CourseMongoServiceSuite) SetupSuite() {
	client, db, repo := InitTestMongoRepo(suite.T())
	serv := service.NewCourse(repo)

	suite.client = client
	suite.db = db
	suite.repo = repo
	suite.service = serv
}

func (suite *CourseMongoServiceSuite) TearDownTest() {
	ctx := context.Background()
	err := suite.db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *CourseMongoServiceSuite) TearDownSuite() {
	ctx := context.Background()
	err := suite.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCourseMongoService(t *testing.T) {
	suite.Run(t, new(CourseMongoServiceSuite))
}

func (suite *CourseMongoServiceSuite) TestGetCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.Get(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *CourseMongoServiceSuite) TestAddCourse() {

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

func (suite *CourseMongoServiceSuite) TestAddLessThanMinStudents() {
	c := factory.NewCourseWithStudents(service.MinNumOfStudents - 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&c)
	suite.Equal(service.InsufficientStudentsErr, err, "Should return InsufficientStudents Err")

}

func (suite *CourseMongoServiceSuite) TestAddMoreThanMaxStudents() {

	c := factory.NewCourseWithStudents(service.MaxNumOfStudents + 1)
	for i := range c.Students {
		err := suite.repo.AddStudent(&c.Students[i])
		suite.Require().Nil(err)
	}
	err := suite.service.Add(&c)
	suite.Require().Equal(service.ExceedingStudentsErr, err, "Should return ExceedingStudents Err")

}

func (suite *CourseMongoServiceSuite) TestEditCourse() {

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

func (suite *CourseMongoServiceSuite) TestEditCourseLessThanMinStudents() {

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

func (suite *CourseMongoServiceSuite) TestEditCourseMoreThanMaxStudents() {

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

func (suite *CourseMongoServiceSuite) TestDeleteCourse() {

	expected := factory.NewCourse()
	err := suite.repo.AddCourse(&expected)
	suite.Require().Nil(err)

	err = suite.service.Delete(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetCourse(expected.Id)
	suite.Require().Equal(err, mongo.ErrNoDocuments)

}
