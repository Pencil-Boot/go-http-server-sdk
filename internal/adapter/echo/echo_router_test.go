package echo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"

	echoLib "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEchoAdapter struct {
	mock.Mock
}

func (m *MockEchoAdapter) Add(method, path string, handler echoLib.HandlerFunc, middleware ...echoLib.MiddlewareFunc) *echoLib.Route {
	args := m.Called(method, path, handler, middleware)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*echoLib.Route)
}

func (m *MockEchoAdapter) Group(prefix string, middleware ...echoLib.MiddlewareFunc) *echoLib.Group {
	args := m.Called(prefix, middleware)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*echoLib.Group)
}

func TestEchoRouter_Register(t *testing.T) {
	mockEcho := new(MockEchoAdapter)
	router := NewEchoRouter(mockEcho)

	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		return nil, nil
	}

	mockEcho.On("Add", "GET", "/test", mock.Anything, mock.Anything).Return(&echoLib.Route{})

	router.Register(httpservercontract.GET, "/test", handler, nil)

	mockEcho.AssertExpectations(t)
}

func TestEchoRouter_Group_Integration(t *testing.T) {
	e := echoLib.New()
	router := NewEchoRouter(e)

	groupRouter := router.Group("/api")

	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		return nil, nil
	}

	groupRouter.Register(httpservercontract.GET, "/users", handler, nil)

	found := false
	for _, route := range e.Routes() {
		if route.Path == "/api/users" && route.Method == "GET" {
			found = true
			break
		}
	}
	assert.True(t, found, "Route should have been registered")
}

func TestConvertHandler(t *testing.T) {
	e := echoLib.New()

	called := false
	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		called = true
		return nil, nil
	}

	echoHandler := convertHandler(handler)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	err := echoHandler(ctx)

	assert.NoError(t, err)
	assert.True(t, called)
}

func TestConvertHandler_Error(t *testing.T) {
	e := echoLib.New()

	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		return nil, assert.AnError
	}

	echoHandler := convertHandler(handler)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	err := echoHandler(ctx)

	assert.Error(t, err)
	he, ok := err.(*echoLib.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, 500, he.Code)
}
