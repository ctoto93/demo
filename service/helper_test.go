package service

import (
	"context"
	"testing"

	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "demo_test"
)

func InitTestMongoRepo(t *testing.T) (*mongo.Client, *mongo.Database, *mongoRepo.Repository) {
	opts := options.Client().ApplyURI(testDbUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		t.Fatal(err)
	}

	db := c.Database(testDbName)
	repo := &mongoRepo.Repository{Db: db}
	return c, db, repo
}
