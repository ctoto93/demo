package rpc_test

import (
	"context"
	"testing"

	"github.com/ctoto93/demo/rpc/pb"
	"github.com/ctoto93/demo/test/factory"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestMongoGRPCStudentFlow(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	clm, db, repo := InitTestMongoRepo(t)
	conn, client := initGRPC(t, repo)

	defer func() {
		db.Drop(ctx)
		clm.Disconnect(ctx)
		conn.Close()
	}()

	s := factory.NewStudent()
	var pbs pb.Student
	err := mapstructure.Decode(s, &pbs)
	require.Nil(err)

	respAddStudent, err := client.AddStudent(ctx, &pbs)
	require.Nil(err)
	require.NotEmpty(respAddStudent.Id)

	respGetStudent, err := client.GetStudent(ctx, &wrappers.StringValue{Value: respAddStudent.Id})
	require.Nil(err)
	checkProtoEqual(require, respAddStudent, respGetStudent)

	editStudent := respGetStudent
	editStudent.Name = "Edit"

	respEditStudent, err := client.EditStudent(ctx, editStudent)
	require.Nil(err)
	checkProtoEqual(require, editStudent, respEditStudent)

	respGetStudent, err = client.GetStudent(ctx, &wrappers.StringValue{Value: respEditStudent.Id})
	require.Nil(err)
	checkProtoEqual(require, respGetStudent, respEditStudent)

	_, err = client.DeleteStudent(ctx, &wrappers.StringValue{Value: respEditStudent.Id})
	require.Nil(err)

	_, err = client.GetStudent(ctx, &wrappers.StringValue{Value: respEditStudent.Id})
	require.NotEmpty(err)

}
