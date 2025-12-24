package httpservercontract

import (
	"context"
	"time"
)

// Server represents the public HTTP server interface.
// Implementations should handle HTTP server lifecycle including starting,
// stopping, and graceful shutdown.
type Server interface {
	// Router returns the router for route registration
	Router() Router

	// Start starts the HTTP server and blocks until the server is stopped
	Start() error

	// StartTLS starts the HTTPS server with the provided certificate and key files
	StartTLS(certFile, keyFile string) error

	// Shutdown gracefully shuts down the server without interrupting active connections
	Shutdown(ctx context.Context) error
}

type ServerConfig struct {
	port           int
	readTimeout    time.Duration
	writeTimeout   time.Duration
	idleTimeout    time.Duration
	maxHeaderBytes int
}

type ServerOption interface {
	applyServer(*ServerConfig)
}

func (c *ServerConfig) Port() int {
	return c.port
}

func (c *ServerConfig) ReadTimeout() time.Duration {
	return c.readTimeout
}

func (c *ServerConfig) WriteTimeout() time.Duration {
	return c.writeTimeout
}

func (c *ServerConfig) IdleTimeout() time.Duration {
	return c.idleTimeout
}

func (c *ServerConfig) MaxHeaderBytes() int {
	return c.maxHeaderBytes
}

func NewServerConfig(opts ...ServerOption) *ServerConfig {
	//TODO: create variables to store default values
	c := &ServerConfig{
		port:           8080,
		readTimeout:    15 * time.Second,
		writeTimeout:   15 * time.Second,
		idleTimeout:    60 * time.Second,
		maxHeaderBytes: 1 << 20, // 1 MB
	}

	for _, opt := range opts {
		opt.applyServer(c)
	}

	return c
}
