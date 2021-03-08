package rpc_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"log"

	mongoRepo "github.com/ctoto93/demo/db/mongo"
	"github.com/ctoto93/demo/rpc"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/ctoto93/demo/service"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "demo_test"
	bufSize    = 1024 * 1024
)

var lis *bufconn.Listener

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

func initGRPC(t *testing.T, repo service.Repository) (*grpc.ClientConn, pb.DemoServiceClient) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterDemoServiceServer(s, rpc.NewDemoService(repo))

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

func checkProtoEqual(r *require.Assertions, expected, actual proto.Message) {
	r.True(
		proto.Equal(expected, actual),
		fmt.Sprintf("These two protobuf messages are not equal:\nexpected: %v\nactual:  %v", expected, actual),
	)
}
