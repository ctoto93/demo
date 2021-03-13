package rpc

import (
	"context"
	"net"
	"testing"

	"log"

	"github.com/ctoto93/demo"
	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/rpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "demo_test"
	bufSize    = 1024 * 1024
)

var lis *bufconn.Listener

func InitTestMongoRepo(t *testing.T) (*mongo.Client, *mongo.Database, demo.Repository) {
	opts := options.Client().ApplyURI(testDbUri)
	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		t.Fatal(err)
	}

	db := c.Database(testDbName)
	repo := mongoRepo.NewRepositoryWithDb(db)
	return c, db, repo
}

func initGRPC(t *testing.T, repo demo.Repository) (*grpc.ClientConn, pb.DemoServiceClient) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	demoRPCServer := NewServer(repo)
	pb.RegisterDemoServiceServer(s, demoRPCServer)

	bufDialer := func(context.Context, string) (net.Conn, error) { return lis.Dial() }

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	client := pb.NewDemoServiceClient(conn)
	if err != nil {
		t.Fatal(err)
	}

	return conn, client

}
