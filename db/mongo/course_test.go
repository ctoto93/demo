package mongo

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (suite *MongoSuite) TestGetCourse() {
	c, err := suite.repo.GetCourse(suite.expectedCourse.Id)
	suite.Require().Nil(err)
	suite.Require().Equal(suite.expectedCourse, c)
}

func (suite *MongoSuite) TestDeleteCourse() {
	err := suite.repo.DeleteCourse(suite.expectedCourse.Id)
	suite.Require().Nil(err)
	count, err := suite.db.Collection("courses").CountDocuments(context.TODO(), bson.M{})
	suite.Require().EqualValues(0, count)
}

func (suite *MongoSuite) TestEditCourse() {
	suite.expectedCourse.Name = "Edited"
	err := suite.repo.EditCourse(&suite.expectedCourse)
	suite.Require().Nil(err)

	var c Course
	oid, err := primitive.ObjectIDFromHex(suite.expectedCourse.Id)
	suite.Require().Nil(err)
	err = suite.db.Collection("courses").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&c)
	suite.Require().Nil(err)
	suite.Require().Equal(c.Name, suite.expectedCourse.Name)

	//check newly added student
	ns := factory.NewStudent()
	err = suite.repo.AddStudent(&ns)
	suite.Require().Nil(err)

	suite.expectedCourse.Students = append(suite.expectedCourse.Students, demo.Student{Id: ns.Id})
	err = suite.repo.EditCourse(&suite.expectedCourse)
	suite.Require().Nil(err)

	ds, err := suite.repo.GetStudent(ns.Id)
	suite.Require().Nil(err)

	suite.Require().True(ds.HasCourse(suite.expectedCourse), "Course should be also in student.courses array")
}

func (suite *MongoSuite) TestAddCourse() {
	dc := factory.NewCourse()

	dc.Students = []demo.Student{
		{Id: suite.expectedStudent.Id},
	}

	err := suite.repo.AddCourse(&dc)
	suite.Require().Nil(err)

	oid, err := primitive.ObjectIDFromHex(dc.Id)
	suite.Require().Nil(err)

	c, err := suite.repo.getCourseByObjectId(oid)
	suite.Require().Nil(err)

	suite.Require().Equal(dc.Id, c.toDemo().Id)
	suite.Require().Equal(dc.Students[0].Id, c.Students[0].Hex())

	ds, err := suite.repo.GetStudent(suite.expectedStudent.Id)
	suite.Require().Nil(err)

	suite.Require().True(ds.HasCourse(dc), "Course should be also in student.courses array")
}
