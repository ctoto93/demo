package integration_test

import (
	"context"
	"log"
	"os"
	"testing"

	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/db/sql"
	"github.com/ctoto93/demo/service"
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

func InitTestSQLiteRepo(t *testing.T) (*gorm.DB, service.Repository) {
	db, err := gorm.Open(sqlite.Open(testSQLitePath), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	tx.AutoMigrate(sql.Course{}, sql.Student{})

	repo := sql.NewRepository(tx)

	return tx, repo
}

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

func buildSQLCleanUpFunc(tx *gorm.DB) func() {
	return func() {
		tx.Rollback()
	}

}

func buildSQLClientDisconnect() func() {
	return func() {
		err := os.Remove(testSQLitePath)
		if err != nil {
			log.Fatal(err)
		}
	}

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
