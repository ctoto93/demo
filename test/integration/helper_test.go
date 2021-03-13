package integration_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func TestMongoService(t *testing.T) {
	suite.Run(t, new(MongoServiceSuite))
}

func (suite *MongoServiceSuite) SetupSuite() {
	host := getEnv("MONGODB_HOST", "localhost")
	port := getEnv("MONGODB_PORT", "27017")

	testDbUri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	opts := options.Client().ApplyURI(testDbUri)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.client = client

}

func (suite *MongoServiceSuite) SetupTest() {

	testDbName := "test_" + strconv.FormatInt(rand.Int63(), 10)
	db := suite.client.Database(testDbName)
	repo := mongoRepo.NewRepositoryWithDb(db)

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
