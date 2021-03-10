package mongo

import (
	"context"
	"fmt"

	"github.com/ctoto93/demo/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Db *mongo.Database
}

func NewRepository(uri string, dbname string) service.Repository {

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

	return &Repository{
		Db: c.Database(dbname),
	}
}
