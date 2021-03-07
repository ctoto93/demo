package rpc

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (service *DemoService) GetCourse(ctx context.Context, id *wrapperspb.StringValue) (*pb.Course, error) {
	course, err := service.course.Get(id.Value)

	if err != nil {
		return nil, err
	}

	var pbCourse pb.Course
	err = mapstructure.Decode(course, &pbCourse)
	if err != nil {
		return nil, err
	}

	return &pbCourse, nil
}

func (service *DemoService) AddCourse(ctx context.Context, newCourse *pb.Course) (*pb.Course, error) {
	var demoCourse demo.Course
	err := mapstructure.Decode(newCourse, &demoCourse)
	if err != nil {
		return nil, err
	}

	err = service.course.Add(&demoCourse)
	if err != nil {
		return nil, err
	}

	newCourse.Id = demoCourse.Id

	return newCourse, nil
}

func (service *DemoService) EditCourse(ctx context.Context, updateCourse *pb.Course) (*pb.Course, error) {
	var demoCourse demo.Course
	err := mapstructure.Decode(updateCourse, &demoCourse)
	if err != nil {
		return nil, err
	}

	err = service.course.Edit(&demoCourse)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(demoCourse, &updateCourse)
	if err != nil {
		return nil, err
	}

	return updateCourse, nil
}

func (service *DemoService) DeleteCourse(ctx context.Context, id *wrapperspb.StringValue) (*emptypb.Empty, error) {
	err := service.course.Delete(id.Value)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
