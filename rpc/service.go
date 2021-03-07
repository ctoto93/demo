package rpc

import (
	"github.com/ctoto93/demo/demopb"
	"github.com/ctoto93/demo/service"
)

type DemoService struct {
	demopb.UnimplementedDemoServiceServer
	course  *service.Course
	student *service.Student
}

func NewDemoService(repo service.Repository) *DemoService {
	return &DemoService{
		course:  service.NewCourse(repo),
		student: service.NewStudent(repo),
	}
}
