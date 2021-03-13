package rpc

import (
	"context"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type GRPCStudentSuite struct {
	suite.Suite
	service *demo.Service
	repo    *mocks.Repository
	server  *server
	conn    *grpc.ClientConn
	client  pb.DemoServiceClient
}

func (suite *GRPCStudentSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	conn, client := initGRPC(suite.T(), suite.repo)
	suite.service = demo.NewService(suite.repo)
	suite.server = NewServer(suite.repo)
	suite.conn = conn
	suite.client = client
}

func (suite *GRPCStudentSuite) TearDownTest() {
	suite.conn.Close()
}

func TestGRPCStudentSuite(t *testing.T) {
	suite.Run(t, new(GRPCStudentSuite))
}

func (suite *GRPCStudentSuite) TestGetStudent() {
	ctx := context.TODO()
	expected := factory.NewStudentWithId()
	suite.repo.On("GetStudent", expected.Id).Return(expected, nil)

	resp, err := suite.client.GetStudent(ctx, &wrappers.StringValue{Value: expected.Id})
	suite.Require().Nil(err)

	actual, err := demo.ToStudent(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCStudentSuite) TestAddStudent() {
	ctx := context.TODO()
	newStudent := factory.NewStudent()
	expected := newStudent

	suite.repo.On("AddStudent", &newStudent).Return(nil)

	var pbs pb.Student
	err := mapstructure.Decode(newStudent, &pbs)
	suite.Require().Nil(err)

	resp, err := suite.client.AddStudent(ctx, &pbs)
	suite.Require().Nil(err)

	actual, err := demo.ToStudent(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCStudentSuite) TestEditStudent() {
	ctx := context.TODO()
	expected := factory.NewStudentWithId()

	suite.repo.On("EditStudent", &expected).Return(nil)

	var pbs pb.Student
	err := mapstructure.Decode(expected, &pbs)
	suite.Require().Nil(err)

	resp, err := suite.client.EditStudent(ctx, &pbs)
	suite.Require().Nil(err)

	actual, err := demo.ToStudent(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCStudentSuite) TestDeleteStudent() {
	ctx := context.TODO()
	expected := factory.NewStudentWithId()

	suite.repo.On("DeleteStudent", expected.Id).Return(nil)

	_, err := suite.client.DeleteStudent(ctx, &wrappers.StringValue{Value: expected.Id})
	suite.Require().Nil(err)

	suite.repo.AssertExpectations(suite.T())
}
