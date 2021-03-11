package rpc

import (
	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
)

type DemoService struct {
	pb.UnimplementedDemoServiceServer
	service *demo.Service
}

func NewDemoService(repo demo.Repository) *DemoService {
	return &DemoService{
		service: demo.NewService(repo),
	}
}
