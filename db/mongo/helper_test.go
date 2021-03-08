package mongo

import (
	"context"
	"flag"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var short = flag.Bool("short", false, "Ignore DB tests")

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "demo_test"
)

type MongoSuite struct {
	suite.Suite
	client          *mongo.Client
	db              *mongo.Database
	repo            *Repository
	expectedStudent demo.Student
	expectedCourse  demo.Course
}

func (suite *MongoSuite) SetupSuite() {
	opts := options.Client().ApplyURI(testDbUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	suite.client = c
}

func (suite *MongoSuite) SetupTest() {
	var err error

	suite.db = suite.client.Database(testDbName)
	suite.repo = &Repository{Db: suite.db}

	c, err := newCourse(factory.NewCourse())
	c.Id = primitive.NewObjectID()
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	s, err := newStudent((factory.NewStudent()))
	if err != nil {
		suite.T().Fatal(err)
		return
	}
	s.Id = primitive.NewObjectID()

	s.Courses = append(s.Courses, c.Id)
	c.Students = append(c.Students, s.Id)

	_, err = suite.db.Collection("courses").InsertOne(context.TODO(), c)
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	_, err = suite.db.Collection("students").InsertOne(context.TODO(), s)
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	suite.expectedStudent = s.toDemo()
	suite.expectedStudent.Courses = []demo.Course{c.toDemo()}
	suite.expectedCourse = c.toDemo()
	suite.expectedCourse.Students = []demo.Student{s.toDemo()}

}

func (suite *MongoSuite) TearDownTest() {
	suite.db.Drop(context.TODO())
}

func (suite *MongoSuite) TearDownSuite() {
	suite.client.Disconnect(context.TODO())
}

func TestMongoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Mongo DB Suite")
		return
	}
	suite.Run(t, new(MongoSuite))
}
