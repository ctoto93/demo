package rpc

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/demopb"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (service *DemoService) GetCourse(ctx context.Context, id *wrapperspb.StringValue) (*demopb.Course, error) {
	course, err := service.course.Get(id.Value)

	if err != nil {
		return nil, err
	}

	var pbCourse demopb.Course
	err = mapstructure.Decode(course, &pbCourse)
	if err != nil {
		return nil, err
	}

	return &pbCourse, nil
}

func (service *DemoService) AddCourse(ctx context.Context, newCourse *demopb.Course) (*demopb.Course, error) {
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

func (service *DemoService) EditCourse(ctx context.Context, updateCourse *demopb.Course) (*demopb.Course, error) {
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
