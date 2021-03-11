package rpc_test

import (
	"context"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/rpc/pb"
	"github.com/ctoto93/demo/test/factory"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func createStudents(t *testing.T, repo demo.Repository, num int) []demo.Student {
	var students []demo.Student
	for i := 0; i < num; i++ {
		s := factory.NewStudent()
		err := repo.AddStudent(&s)
		if err != nil {
			t.Fatal(err)
		}
		students = append(students, s)
	}
	return students
}

func TestMongoGRPCCourseFlow(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	clm, db, repo := InitTestMongoRepo(t)
	conn, client := initGRPC(t, repo)

	defer func() {
		db.Drop(ctx)
		clm.Disconnect(ctx)
		conn.Close()
	}()

	c := factory.NewCourse()
	students := createStudents(t, repo, demo.MinNumOfStudents)
	c.Students = students

	var pbs pb.Course
	err := mapstructure.Decode(c, &pbs)
	require.Nil(err)

	respAddCourse, err := client.AddCourse(ctx, &pbs)
	require.Nil(err)
	require.NotEmpty(respAddCourse.Id)

	respGetCourse, err := client.GetCourse(ctx, &wrappers.StringValue{Value: respAddCourse.Id})
	require.Nil(err)
	checkProtoEqual(require, respAddCourse, respGetCourse)

	editCourse := respGetCourse
	editCourse.Name = "Edit"

	respEditCourse, err := client.EditCourse(ctx, editCourse)
	require.Nil(err)
	checkProtoEqual(require, editCourse, respEditCourse)

	respGetCourse, err = client.GetCourse(ctx, &wrappers.StringValue{Value: respEditCourse.Id})
	require.Nil(err)
	checkProtoEqual(require, respGetCourse, respEditCourse)

	_, err = client.DeleteCourse(ctx, &wrappers.StringValue{Value: respEditCourse.Id})
	require.Nil(err)

	_, err = client.GetCourse(ctx, &wrappers.StringValue{Value: respEditCourse.Id})
	require.NotEmpty(err)

}
