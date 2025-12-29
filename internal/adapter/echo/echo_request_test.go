package echo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EchoRequestSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestEchoRequestSuite(t *testing.T) {
	suite.Run(t, new(EchoRequestSuite))
}

func (s *EchoRequestSuite) SetupTest() {
	s.echo = echo.New()
}

func (s *EchoRequestSuite) TestRequest_BasicInfo() {
	request := httptest.NewRequest(http.MethodPost, "/test?q=1", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	r := newEchoRequest(ctx)

	assert.Equal(s.T(), http.MethodPost, r.Method())
	assert.Equal(s.T(), "/test", r.Path())
}

func (s *EchoRequestSuite) TestRequest_Headers() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Add("X-Test", "value1")
	ctx := s.echo.NewContext(request, nil)

	r := newEchoRequest(ctx)

	assert.Equal(s.T(), "value1", r.Header("X-Test"))
	assert.Equal(s.T(), []string{"value1"}, r.Headers().Get("X-Test"))

	assert.NotNil(s.T(), r.Headers())
	assert.Equal(s.T(), map[string][]string{"X-Test": {"value1"}}, r.Headers().All())
}

func (s *EchoRequestSuite) TestRequest_Params() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := s.echo.NewContext(request, nil)
	ctx.SetParamNames("id")
	ctx.SetParamValues("123")

	r := newEchoRequest(ctx)

	assert.Equal(s.T(), "123", r.Param("id"))
	assert.Equal(s.T(), map[string]string{"id": "123"}, r.Params().All())

	assert.NotNil(s.T(), r.Params())
	assert.Equal(s.T(), "123", r.Params().Get("id"))
}

func (s *EchoRequestSuite) TestRequest_QueryParams() {
	request := httptest.NewRequest(http.MethodGet, "/?page=1", nil)
	ctx := s.echo.NewContext(request, nil)

	r := newEchoRequest(ctx)

	assert.Equal(s.T(), "1", r.QueryParam("page"))
	assert.Equal(s.T(), map[string][]string{"page": {"1"}}, r.QueryParams().All())

	assert.NotNil(s.T(), r.QueryParams())
	assert.Equal(s.T(), []string{"1"}, r.QueryParams().Get("page"))
}

func (s *EchoRequestSuite) TestRequest_Bind() {
	body := `{"name":"test"}`
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := s.echo.NewContext(request, nil)

	r := newEchoRequest(ctx)

	var target struct {
		Name string `json:"name"`
	}
	err := r.Bind(&target)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "test", target.Name)
}

func (s *EchoRequestSuite) TestRequest_BindError() {
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("invalid json"))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := s.echo.NewContext(request, nil)

	r := newEchoRequest(ctx)

	var target struct {
		Name string `json:"name"`
	}
	err := r.Bind(&target)
	assert.Error(s.T(), err)
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, assert.AnError
}

func (s *EchoRequestSuite) TestRequest_BindReadError() {
	request := httptest.NewRequest(http.MethodPost, "/", &errorReader{})
	ctx := s.echo.NewContext(request, nil)

	r := newEchoRequest(ctx)

	var target struct{}
	err := r.Bind(&target)
	assert.Error(s.T(), err)
}
