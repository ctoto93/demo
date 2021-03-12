package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctoto93/demo"
	"github.com/ctoto93/demo/test/factory"
	"github.com/ctoto93/demo/test/mocks"
	"github.com/stretchr/testify/suite"
)

type RESTStudentSuite struct {
	suite.Suite
	service *demo.Service
	repo    *mocks.Repository
	server  *httptest.Server
}

func TestRESTStudent(t *testing.T) {
	suite.Run(t, new(RESTStudentSuite))

}

func (suite *RESTStudentSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.service = demo.NewService(suite.repo)
	server := NewServer(suite.repo)
	suite.server = httptest.NewServer(server.router)
}

func (suite *RESTStudentSuite) TearDownTest() {
	suite.server.Close()
}

func parseBody(r io.Reader, out *Response) error {
	if err := json.NewDecoder(r).Decode(out); err != nil {
		return err
	}
	return nil
}

func (suite *RESTStudentSuite) TestGetStudent() {
	expected := factory.NewStudentWithId()
	suite.repo.On("GetStudent", expected.Id).Return(expected, nil)

	httpResp, err := http.Get(fmt.Sprintf("%s/v1/students/%s", suite.server.URL, expected.Id))
	suite.Require().Nil(err)
	defer httpResp.Body.Close()

	suite.Require().Equal(http.StatusOK, httpResp.StatusCode)

	var resp Response
	err = parseBody(httpResp.Body, &resp)
	suite.Require().Nil(err)

	actual, err := demo.ToStudent(resp.Data)
	suite.Require().Nil(err)

	suite.Require().Equal(expected, actual)

}
