package httpservercontract

import (
	"context"
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
