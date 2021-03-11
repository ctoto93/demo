package integration_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ctoto93/demo"
	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/db/sql"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	testDbUri      = "mongodb://localhost:27017"
	testDbName     = "demo_test"
	testSQLitePath = "gorm.db"
)

type ServiceSuite struct {
	suite.Suite
	repo    demo.Repository
	service *demo.Service
}

type MongoServiceSuite struct {
	ServiceSuite
	client *mongo.Client
	db     *mongo.Database
}

type SQLiteServiceSuite struct {
	ServiceSuite
	db *gorm.DB
	tx *gorm.DB
}

func TestMongoService(t *testing.T) {
	suite.Run(t, new(MongoServiceSuite))
}

func (suite *MongoServiceSuite) SetupSuite() {
	opts := options.Client().ApplyURI(testDbUri)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		suite.T().Fatal(err)
	}

	db := client.Database(testDbName)
	repo := mongoRepo.NewRepositoryWithDb(db)
	suite.client = client
	suite.db = db
	suite.repo = repo
	suite.service = demo.NewService(repo)
}

func (suite *MongoServiceSuite) TearDownTest() {
	ctx := context.Background()
	err := suite.db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *MongoServiceSuite) TearDownSuite() {
	ctx := context.Background()
	err := suite.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSQLiteService(t *testing.T) {
	suite.Run(t, new(SQLiteServiceSuite))
}

func (suite *SQLiteServiceSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(testSQLitePath), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}

	if err := db.AutoMigrate(sql.Course{}, sql.Student{}); err != nil {
		suite.T().Fatal()
	}

	suite.db = db
}

func (suite *SQLiteServiceSuite) SetupTest() {
	suite.tx = suite.db.Begin()
	suite.repo = sql.NewRepository(suite.tx)
	suite.service = demo.NewService(suite.repo)
}

func (suite *SQLiteServiceSuite) TearDownTest() {
	suite.tx.Rollback()
}

func (suite *SQLiteServiceSuite) TearDownSuite() {
	err := os.Remove(testSQLitePath)
	if err != nil {
		log.Fatal(err)
	}
}
