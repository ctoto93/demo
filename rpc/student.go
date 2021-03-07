package rpc

import (
	"context"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (service *DemoService) GetStudent(ctx context.Context, id *wrapperspb.StringValue) (*pb.Student, error) {
	student, err := service.student.Get(id.Value)

	if err != nil {
		return nil, err
	}

	var pbStudent pb.Student
	err = mapstructure.Decode(student, &pbStudent)
	if err != nil {
		return nil, err
	}

	return &pbStudent, nil
}

func (service *DemoService) AddStudent(ctx context.Context, newStudent *pb.Student) (*pb.Student, error) {
	var demoStudent demo.Student
	err := mapstructure.Decode(newStudent, &demoStudent)
	if err != nil {
		return nil, err
	}

	err = service.student.Add(&demoStudent)
	if err != nil {
		return nil, err
	}

	newStudent.Id = demoStudent.Id

	return newStudent, nil
}

func (service *DemoService) EditStudent(ctx context.Context, updateStudent *pb.Student) (*pb.Student, error) {
	var demoStudent demo.Student
	err := mapstructure.Decode(updateStudent, &demoStudent)
	if err != nil {
		return nil, err
	}

	err = service.student.Edit(&demoStudent)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(demoStudent, &updateStudent)
	if err != nil {
		return nil, err
	}

	return updateStudent, nil
}

func (service *DemoService) DeleteStudent(ctx context.Context, id *wrapperspb.StringValue) (*emptypb.Empty, error) {
	err := service.student.Delete(id.Value)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
