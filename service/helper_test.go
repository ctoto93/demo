package service_test

import (
	"context"
	"log"
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
	repo := mongoRepo.NewRepositoryWithDb(db)
	return c, db, repo
}

func buildMongoCleanUpFunc(db *mongo.Database) func() {
	return func() {
		ctx := context.Background()
		err := db.Drop(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func buildMongoClientDisconnectFunc(client *mongo.Client) func() {
	return func() {
		ctx := context.Background()
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}
