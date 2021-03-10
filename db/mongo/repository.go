package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	db *mongo.Database
}

func NewRepository(uri string, dbname string) *repository {

	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = c.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &repository{
		db: c.Database(dbname),
	}
}

func NewRepositoryWithDb(db *mongo.Database) *repository {
	return &repository{db: db}
}
