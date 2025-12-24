package echo

import (
	"context"
	"fmt"
	"go-http-server-sdk/internal/core/server"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"net/http"

	echoLib "github.com/labstack/echo/v4"
)

// Server represents the public HTTP server
type echoServer struct {
	echo *echoLib.Echo
}

// NewServer creates a new server
func NewServer(echo *echoLib.Echo) server.ServerExecutor {

	return &echoServer{
		echo: echo,
	}
}

// Start starts the HTTP server
func (s *echoServer) Start(cfg *httpservercontract.ServerConfig) error {
	addr := getAddress(cfg)
	fillServerConfig(cfg, s.echo)

	if err := s.echo.Start(addr); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// StartTLS starts the HTTPS server
func (s *echoServer) StartTLS(certFile, keyFile string, cfg *httpservercontract.ServerConfig) error {
	addr := getAddress(cfg)
	fillServerConfig(cfg, s.echo)

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

func getAddress(cfg *httpservercontract.ServerConfig) string {
	return fmt.Sprintf(":%d", cfg.Port())
}

func fillServerConfig(cfg *httpservercontract.ServerConfig, s *echoLib.Echo) {
	s.Server.ReadTimeout = cfg.ReadTimeout()
	s.Server.WriteTimeout = cfg.WriteTimeout()
	s.Server.IdleTimeout = cfg.IdleTimeout()
	s.Server.MaxHeaderBytes = cfg.MaxHeaderBytes()
}
