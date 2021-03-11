package rpc

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *DemoService) GetCourse(ctx context.Context, id *wrapperspb.StringValue) (*pb.Course, error) {
	course, err := s.service.GetCourse(id.Value)

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

func (s *DemoService) AddCourse(ctx context.Context, newCourse *pb.Course) (*pb.Course, error) {
	var demoCourse demo.Course
	err := mapstructure.Decode(newCourse, &demoCourse)
	if err != nil {
		return nil, err
	}

	err = s.service.AddCourse(&demoCourse)
	if err != nil {
		return nil, err
	}

	newCourse.Id = demoCourse.Id

	return newCourse, nil
}

func (s *DemoService) EditCourse(ctx context.Context, updateCourse *pb.Course) (*pb.Course, error) {
	var demoCourse demo.Course
	err := mapstructure.Decode(updateCourse, &demoCourse)
	if err != nil {
		return nil, err
	}

	err = s.service.EditCourse(&demoCourse)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(demoCourse, &updateCourse)
	if err != nil {
		return nil, err
	}

	return updateCourse, nil
}

func (s *DemoService) DeleteCourse(ctx context.Context, id *wrapperspb.StringValue) (*emptypb.Empty, error) {
	err := s.service.DeleteCourse(id.Value)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
