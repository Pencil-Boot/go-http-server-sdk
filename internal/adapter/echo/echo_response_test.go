package echo

import (
	"encoding/json"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockResponse struct {
	mock.Mock
}

func (m *MockResponse) Status() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockResponse) Data() any {
	args := m.Called()
	return args.Get(0)
}

func (m *MockResponse) Body() any {
	args := m.Called()
	return args.Get(0)
}

func (m *MockResponse) Headers() httpservercontract.ResponseHeaders {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(httpservercontract.ResponseHeaders)
}

type MockResponseHeaders struct {
	mock.Mock
}

func (m *MockResponseHeaders) Get(key string) []string {
	args := m.Called(key)
	return args.Get(0).([]string)
}

func (m *MockResponseHeaders) Set(key, value string) {
	m.Called(key, value)
}

func (m *MockResponseHeaders) Add(key, value string) {
	m.Called(key, value)
}

func (m *MockResponseHeaders) All() map[string][]string {
	args := m.Called()
	return args.Get(0).(map[string][]string)
}

type EchoResponseSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestEchoResponseSuite(t *testing.T) {
	suite.Run(t, new(EchoResponseSuite))
}

func (s *EchoResponseSuite) SetupTest() {
	s.echo = echo.New()
}

func (s *EchoResponseSuite) TestWriteResponse_Nil() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	err := writeResponse(ctx, nil)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, recorder.Code)
}

func (s *EchoResponseSuite) TestWriteResponse_String() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	mockResp := new(MockResponse)
	mockResp.On("Status").Return(201)
	mockResp.On("Body").Return("created")
	mockResp.On("Headers").Return(nil)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 201, recorder.Code)
	assert.Equal(s.T(), "created", recorder.Body.String())
}

func (s *EchoResponseSuite) TestWriteResponse_JSON() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	data := map[string]string{"foo": "bar"}
	mockResp := new(MockResponse)
	mockResp.On("Status").Return(200)
	mockResp.On("Body").Return(data)
	mockResp.On("Headers").Return(nil)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 200, recorder.Code)

	var result map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Equal(s.T(), data, result)
}

func (s *EchoResponseSuite) TestWriteResponse_Headers() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	mockHeaders := new(MockResponseHeaders)
	mockHeaders.On("All").Return(map[string][]string{"X-Foo": {"bar"}})

	mockResp := new(MockResponse)
	mockResp.On("Status").Return(200)
	mockResp.On("Body").Return(nil)
	mockResp.On("Headers").Return(mockHeaders)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "bar", recorder.Header().Get("X-Foo"))
}

func (s *EchoResponseSuite) TestWriteResponse_MarshalError() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	// Channels cannot be marshaled to JSON
	mockResp := new(MockResponse)
	mockResp.On("Status").Return(200)
	mockResp.On("Body").Return(make(chan int))
	mockResp.On("Headers").Return(nil)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 500, recorder.Code)
	assert.Equal(s.T(), "failed to marshal response", recorder.Body.String())
}

func (s *EchoResponseSuite) TestWriteResponse_ByteSlice() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	data := []byte("binary data")
	mockResp := new(MockResponse)
	mockResp.On("Status").Return(200)
	mockResp.On("Body").Return(data)
	mockResp.On("Headers").Return(nil)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 200, recorder.Code)
	assert.Equal(s.T(), data, recorder.Body.Bytes())
	assert.Equal(s.T(), "application/octet-stream", recorder.Header().Get(echo.HeaderContentType))
}

func (s *EchoResponseSuite) TestWriteResponse_StatusZero() {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := s.echo.NewContext(request, recorder)

	mockResp := new(MockResponse)
	mockResp.On("Status").Return(0)
	mockResp.On("Body").Return("ok")
	mockResp.On("Headers").Return(nil)

	err := writeResponse(ctx, mockResp)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 200, recorder.Code)
}
