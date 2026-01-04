package server

import (
	"context"

	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

type server struct {
	router   httpservercontract.Router
	config   *httpservercontract.ServerConfig
	executor ServerExecutor
}

func NewServer(
	router httpservercontract.Router,
	config *httpservercontract.ServerConfig,
	executor ServerExecutor,
) httpservercontract.Server {
	return &server{
		router:   router,
		config:   config,
		executor: executor,
	}
}

// Router returns the router for route registration
func (s *server) Router() httpservercontract.Router {
	return s.router
}

// Start starts the server
func (s *server) Start() error {
	return s.executor.Start(s.config)
}

// StartTLS starts the server with TLS
func (s *server) StartTLS(certFile, keyFile string) error {
	return s.executor.StartTLS(certFile, keyFile, s.config)
}

// Shutdown gracefully shuts down the server
func (s *server) Shutdown(ctx context.Context) error {
	return s.executor.Shutdown(ctx)
}
