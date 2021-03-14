package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/stretchr/testify/suite"
)

type RESTCourseSuite struct {
	suite.Suite
	service *demo.Service
	repo    *mocks.Repository
	server  *httptest.Server
}

func TestRESTCourse(t *testing.T) {
	suite.Run(t, new(RESTCourseSuite))

}

func (suite *RESTCourseSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.service = demo.NewService(suite.repo)
	server := NewServer(suite.repo)
	suite.server = httptest.NewServer(server.router)
}

func (suite *RESTCourseSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *RESTCourseSuite) TestGetCourse() {
	expected := factory.NewCourseWithId()
	suite.repo.On("GetCourse", expected.Id).Return(expected, nil)

	httpResp, err := http.Get(fmt.Sprintf("%s/v1/courses/%s", suite.server.URL, expected.Id))
	suite.Require().Nil(err)
	defer httpResp.Body.Close()

	suite.Require().Equal(http.StatusOK, httpResp.StatusCode)

	var resp Response
	err = parseBody(httpResp.Body, &resp)
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp.Data)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *RESTCourseSuite) TestAddCourse() {
	newCourse := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	expected := newCourse

	suite.repo.On("AddCourse", &newCourse).Once().Return(nil)

	req, err := json.Marshal(newCourse)
	suite.Require().Nil(err)

	httpResp, err := http.Post(fmt.Sprintf("%s/v1/courses/", suite.server.URL), "application/JSON", bytes.NewBuffer(req))
	suite.Require().Nil(err)
	defer httpResp.Body.Close()

	suite.Require().Equal(http.StatusOK, httpResp.StatusCode)

	var resp Response
	bytes, _ := httputil.DumpResponse(httpResp, true)
	fmt.Println(string(bytes))
	err = parseBody(httpResp.Body, &resp)
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp.Data)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *RESTCourseSuite) TestEditCourse() {
	expected := factory.NewCourseWithStudents(demo.MinNumOfStudents)
	expected.Id = "1"

	suite.repo.On("EditCourse", &expected).Once().Return(nil)

	body, err := json.Marshal(expected)
	suite.Require().Nil(err)

	url := fmt.Sprintf("%s/v1/courses/%s", suite.server.URL, expected.Id)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	suite.Require().Nil(err)
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := http.DefaultClient.Do(req)
	suite.Require().Nil(err)
	defer httpResp.Body.Close()

	suite.Require().Equal(http.StatusOK, httpResp.StatusCode)

	var resp Response
	err = parseBody(httpResp.Body, &resp)
	suite.Require().Nil(err)

	actual, err := demo.ToCourse(resp.Data)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *RESTCourseSuite) TestDeleteCourse() {
	expected := factory.NewCourseWithId()

	suite.repo.On("DeleteCourse", expected.Id).Once().Return(nil)

	url := fmt.Sprintf("%s/v1/courses/%s", suite.server.URL, expected.Id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	suite.Require().Nil(err)
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := http.DefaultClient.Do(req)
	suite.Require().Nil(err)
	defer httpResp.Body.Close()

	suite.Require().Equal(http.StatusOK, httpResp.StatusCode)

	var resp Response
	err = parseBody(httpResp.Body, &resp)
	suite.Require().Nil(err)

	suite.Require().True(resp.Meta.Success)
	suite.repo.AssertExpectations(suite.T())
}
