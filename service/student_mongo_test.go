package service

import (
	"context"
	"log"
	"testing"

	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/test/factory"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentMongoServiceSuite struct {
	suite.Suite
	client  *mongo.Client
	db      *mongo.Database
	repo    *mongoRepo.Repository
	service *Student
}

func (suite *StudentMongoServiceSuite) SetupSuite() {
	client, db, repo := InitTestMongoRepo(suite.T())
	serv := NewStudent(repo)

	suite.client = client
	suite.db = db
	suite.repo = repo
	suite.service = serv
}

func (suite *StudentMongoServiceSuite) TearDownTest() {
	ctx := context.Background()
	err := suite.db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *StudentMongoServiceSuite) TearDownSuite() {
	ctx := context.Background()
	err := suite.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStudentMongoService(t *testing.T) {
	suite.Run(t, new(StudentMongoServiceSuite))
}

func (suite *StudentMongoServiceSuite) TestGetStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	actual, err := suite.service.Get(expected.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(expected, actual)

}

func (suite *StudentMongoServiceSuite) TestAddStudent() {

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

func (suite *StudentMongoServiceSuite) TestAddStudentExceedingCreditLimit() {
	expected := factory.NewStudentWithCourses(1)
	expected.Courses[0].Credit = 40
	for i := range expected.Courses {
		err := suite.repo.AddCourse(&expected.Courses[i])
		suite.Require().Nil(err)
	}

	err := suite.service.Add(&expected)
	suite.Require().Equal(OverCreditErr, err, "Should return OverCreditErr")

}

func (suite *StudentMongoServiceSuite) TestEditStudent() {

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

func (suite *StudentMongoServiceSuite) TestDeleteStudent() {

	expected := factory.NewStudent()
	err := suite.repo.AddStudent(&expected)
	suite.Require().Nil(err)

	err = suite.service.Delete(expected.Id)
	suite.Require().Nil(err)

	_, err = suite.repo.GetStudent(expected.Id)
	suite.Require().Equal(err, mongo.ErrNoDocuments)

}
