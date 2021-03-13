package rpc

import (
	"fmt"
	"net"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDemoServiceServer
	service    *demo.Service
	grpcServer *grpc.Server
}

func NewServer(repo demo.Repository) *server {

	return &server{
		service:    demo.NewService(repo),
		grpcServer: grpc.NewServer(),
	}
}

func (s *server) Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	pb.RegisterDemoServiceServer(s.grpcServer, s)
	fmt.Printf("Serving on %s\n", address)
	return s.grpcServer.Serve(lis)

}
