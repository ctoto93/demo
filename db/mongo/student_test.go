package mongo

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (suite *MongoSuite) TestGetStudent() {
	s, err := suite.repo.GetStudent(suite.expectedStudent.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(suite.expectedStudent, s)
}

func (suite *MongoSuite) TestDeleteStudent() {
	err := suite.repo.DeleteStudent(suite.expectedStudent.Id)
	suite.Require().Nil(err)
	count, err := suite.db.Collection("students").CountDocuments(context.TODO(), bson.M{})
	suite.Require().EqualValues(0, count)
}

func (suite *MongoSuite) TestEditStudent() {
	suite.expectedStudent.Name = "Edited"
	err := suite.repo.EditStudent(&suite.expectedStudent)
	suite.Require().Nil(err)

	var s Student
	oid, err := primitive.ObjectIDFromHex(suite.expectedStudent.Id)
	suite.Require().Nil(err)
	err = suite.db.Collection("students").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&s)
	suite.Require().Nil(err)
	suite.Require().Equal(s.Name, suite.expectedStudent.Name)
}

func (suite *MongoSuite) TestAddStudent() {
	ds := factory.NewStudent()

	ds.Courses = []demo.Course{
		{Id: suite.expectedCourse.Id},
	}

	err := suite.repo.AddStudent(&ds)
	suite.Require().Nil(err)

	oid, err := primitive.ObjectIDFromHex(ds.Id)
	suite.Require().Nil(err)

	s, err := suite.repo.getStudentByObjectId(oid)
	suite.Require().Nil(err)

	suite.Require().Equal(ds.Id, s.toDemo().Id)
	suite.Require().Equal(ds.Courses[0].Id, s.Courses[0].Hex())
}
