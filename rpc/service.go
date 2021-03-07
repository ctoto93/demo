package rpc

import (
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/ctoto93/demo/service"
)

type DemoService struct {
	pb.UnimplementedDemoServiceServer
	course  *service.Course
	student *service.Student
}

func NewDemoService(repo service.Repository) *DemoService {
	return &DemoService{
		course:  service.NewCourse(repo),
		student: service.NewStudent(repo),
	}
}
