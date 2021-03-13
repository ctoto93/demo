package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type repository struct {
	db *mongo.Database
}

func NewRepository(uri string) *repository {

	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		panic(err)
	}
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	err = c.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &repository{
		db: c.Database(cs.Database),
	}
}

func NewRepositoryWithDb(db *mongo.Database) *repository {
	return &repository{db: db}
}
