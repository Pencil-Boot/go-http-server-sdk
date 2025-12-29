package router

import (
	"context"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRouterRegister struct {
	mock.Mock
}

func (m *MockRouterRegister) Register(method httpservercontract.HttpMethod, path string, handler httpservercontract.Handler, config *httpservercontract.RouterConfig) {
	m.Called(method, path, handler, config)
}

func (m *MockRouterRegister) Group(prefix string) RouterRegister {
	args := m.Called(prefix)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(RouterRegister)
}

type MockMiddleware struct {
	mock.Mock
}

func (m *MockMiddleware) Handle(ctx context.Context, req httpservercontract.Request, next httpservercontract.Handler) (httpservercontract.Response, error) {
	args := m.Called(ctx, req, next)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(httpservercontract.Response), args.Error(1)
}

func TestRouter_Methods(t *testing.T) {
	mockRegister := new(MockRouterRegister)
	r := NewRouter(mockRegister)

	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		return nil, nil
	}

	methods := []struct {
		name   string
		method func(string, httpservercontract.Handler, ...httpservercontract.RouterOption)
		verb   httpservercontract.HttpMethod
	}{
		{"GET", r.GET, httpservercontract.GET},
		{"POST", r.POST, httpservercontract.POST},
		{"PUT", r.PUT, httpservercontract.PUT},
		{"DELETE", r.DELETE, httpservercontract.DELETE},
		{"PATCH", r.PATCH, httpservercontract.PATCH},
		{"OPTIONS", r.OPTIONS, httpservercontract.OPTIONS},
		{"HEAD", r.HEAD, httpservercontract.HEAD},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			mockRegister.On("Register", tt.verb, "/test", mock.Anything, mock.Anything).Return()
			tt.method("/test", handler)
			mockRegister.AssertExpectations(t)
		})
	}
}

func TestRouter_Group(t *testing.T) {
	mockRegister := new(MockRouterRegister)
	mockGroupRegister := new(MockRouterRegister)
	r := NewRouter(mockRegister)

	mockRegister.On("Group", "/api").Return(mockGroupRegister)

	group := r.Group("/api")
	assert.NotNil(t, group)

	mockRegister.AssertExpectations(t)
}

func TestRouter_Middleware(t *testing.T) {
	mockRegister := new(MockRouterRegister)
	r := NewRouter(mockRegister)

	mockMW := new(MockMiddleware)
	r.Use(mockMW)

	handler := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
		return nil, nil
	}

	mockRegister.On("Register", httpservercontract.GET, "/test", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		h := args.Get(2).(httpservercontract.Handler)
		mockMW.On("Handle", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		h(context.Background(), nil)
	}).Return()

	r.GET("/test", handler)

	mockRegister.AssertExpectations(t)
	mockMW.AssertExpectations(t)
}
