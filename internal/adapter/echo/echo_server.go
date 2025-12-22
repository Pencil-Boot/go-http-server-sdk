package echo

import (
	"context"
	"fmt"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"
	"net/http"

	echoLib "github.com/labstack/echo/v4"
)

// Server represents the public HTTP server
type echoServer struct {
	port   int
	router httpservercontract.Router
	echo   *echoLib.Echo
	config httpservercontract.Config
}

// NewServer creates a new server
func NewServer(port int, cfg httpservercontract.Config) httpservercontract.Server {
	router := NewEchoRouter()
	// Get the Echo instance from the router
	echoRouter, ok := router.(*echoRouter)
	if !ok {
		panic("router is not an echoRouter")
	}

	return &echoServer{
		port:   port,
		router: router,
		echo:   echoRouter.GetEcho(),
		config: cfg,
	}
}

// Router returns the router for route registration
func (s *echoServer) Router() httpservercontract.Router {
	return s.router
}

// Start starts the HTTP server
func (s *echoServer) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.echo.Server.Addr = addr
	s.echo.Server.ReadTimeout = s.config.ReadTimeout()
	s.echo.Server.WriteTimeout = s.config.WriteTimeout()
	s.echo.Server.IdleTimeout = s.config.IdleTimeout()
	s.echo.Server.MaxHeaderBytes = s.config.MaxHeaderBytes()

	if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// StartTLS starts the HTTPS server
func (s *echoServer) StartTLS(certFile, keyFile string) error {
	addr := fmt.Sprintf(":%d", s.port)
	s.echo.Server.Addr = addr
	s.echo.Server.ReadTimeout = s.config.ReadTimeout()
	s.echo.Server.WriteTimeout = s.config.WriteTimeout()
	s.echo.Server.IdleTimeout = s.config.IdleTimeout()
	s.echo.Server.MaxHeaderBytes = s.config.MaxHeaderBytes()

	if err := s.echo.StartTLS(addr, certFile, keyFile); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start TLS server: %w", err)
	}
	return nil
}

// Shutdown gracefully shuts down the server
func (s *echoServer) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	return nil
}
