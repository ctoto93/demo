package rpc

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/api/rpc/pb"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetCourse(ctx context.Context, id *wrapperspb.StringValue) (*pb.Course, error) {
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

func (s *server) AddCourse(ctx context.Context, newCourse *pb.Course) (*pb.Course, error) {
	demoCourse, err := demo.ToCourse(newCourse)
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

func (s *server) EditCourse(ctx context.Context, updateCourse *pb.Course) (*pb.Course, error) {
	demoCourse, err := demo.ToCourse(updateCourse)
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

func (s *server) DeleteCourse(ctx context.Context, id *wrapperspb.StringValue) (*emptypb.Empty, error) {
	err := s.service.DeleteCourse(id.Value)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
