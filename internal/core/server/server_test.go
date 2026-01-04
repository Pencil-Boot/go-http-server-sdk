package server

import (
	"context"
	"testing"

	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServerExecutor struct {
	mock.Mock
}

func (m *MockServerExecutor) Start(cfg *httpservercontract.ServerConfig) error {
	args := m.Called(cfg)
	return args.Error(0)
}

func (m *MockServerExecutor) StartTLS(certFile, keyFile string, cfg *httpservercontract.ServerConfig) error {
	args := m.Called(certFile, keyFile, cfg)
	return args.Error(0)
}

func (m *MockServerExecutor) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

type MockRouter struct {
	mock.Mock
}

func (m *MockRouter) GET(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) POST(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) PUT(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) DELETE(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) PATCH(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) OPTIONS(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) HEAD(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	m.Called(path, handler, opts)
}
func (m *MockRouter) Group(prefix string) httpservercontract.Router {
	args := m.Called(prefix)
	return args.Get(0).(httpservercontract.Router)
}
func (m *MockRouter) Use(middleware ...httpservercontract.Middleware) {
	m.Called(middleware)
}

func TestServer_Lifecycle(t *testing.T) {
	mockRouter := new(MockRouter)
	mockExecutor := new(MockServerExecutor)
	cfg := httpservercontract.NewServerConfig()
	s := NewServer(mockRouter, cfg, mockExecutor)

	t.Run("Router", func(t *testing.T) {
		assert.Equal(t, mockRouter, s.Router())
	})

	t.Run("Start", func(t *testing.T) {
		mockExecutor.On("Start", cfg).Return(nil)
		err := s.Start()
		assert.NoError(t, err)
	})

	t.Run("StartTLS", func(t *testing.T) {
		mockExecutor.On("StartTLS", "cert", "key", cfg).Return(nil)
		err := s.StartTLS("cert", "key")
		assert.NoError(t, err)
	})

	t.Run("Shutdown", func(t *testing.T) {
		ctx := context.Background()
		mockExecutor.On("Shutdown", ctx).Return(nil)
		err := s.Shutdown(ctx)
		assert.NoError(t, err)
	})

	mockExecutor.AssertExpectations(t)
}
