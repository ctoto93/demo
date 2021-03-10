package service_test

import (
	"context"
	"testing"

	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "demo_test"
)

func InitTestMongoRepo(t *testing.T) (*mongo.Client, *mongo.Database, service.Repository) {
	opts := options.Client().ApplyURI(testDbUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		t.Fatal(err)
	}

	db := c.Database(testDbName)
	repo := &mongoRepo.NewRepositoryWithDb(db)
	return c, db, repo
}
