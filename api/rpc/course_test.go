package rpc

import (
	"context"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/api/rpc/pb"
	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type GRPCCourseSuite struct {
	suite.Suite
	service *demo.Service
	repo    *mocks.Repository
	server  *server
	conn    *grpc.ClientConn
	client  pb.DemoServiceClient
}

func (suite *GRPCCourseSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	conn, client := initGRPC(suite.T(), suite.repo)
	suite.service = demo.NewService(suite.repo)
	suite.server = NewServer(suite.repo)
	suite.conn = conn
	suite.client = client
}

func (suite *GRPCCourseSuite) TearDownTest() {
	suite.conn.Close()
}

func TestGRPCCourseSuite(t *testing.T) {
	suite.Run(t, new(GRPCCourseSuite))
}

func (suite *GRPCCourseSuite) TestGetCourse() {
	ctx := context.TODO()
	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	suite.repo.On("GetCourse", expected.Id).Return(expected, nil)

	resp, err := suite.client.GetCourse(ctx, &wrappers.StringValue{Value: expected.Id})
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCCourseSuite) TestAddCourse() {
	ctx := context.TODO()
	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)

	suite.repo.On("AddCourse", &expected).Return(nil)

	var pbs pb.Course
	err := mapstructure.Decode(expected, &pbs)
	suite.Require().Nil(err)

	resp, err := suite.client.AddCourse(ctx, &pbs)
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCCourseSuite) TestEditCourse() {
	ctx := context.TODO()
	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)

	suite.repo.On("EditCourse", &expected).Return(nil)

	var pbs pb.Course
	err := mapstructure.Decode(expected, &pbs)
	suite.Require().Nil(err)

	resp, err := suite.client.EditCourse(ctx, &pbs)
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *GRPCCourseSuite) TestDeleteCourse() {
	ctx := context.TODO()
	expected := factory.NewCourseWithId()

	suite.repo.On("DeleteCourse", expected.Id).Return(nil)

	_, err := suite.client.DeleteCourse(ctx, &wrappers.StringValue{Value: expected.Id})
	suite.Require().Nil(err)

	suite.repo.AssertExpectations(suite.T())
}
