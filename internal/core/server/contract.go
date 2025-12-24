package server

import (
	"context"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

type ServerExecutor interface {
	Start(config *httpservercontract.ServerConfig) error
	StartTLS(certFile, keyFile string, config *httpservercontract.ServerConfig) error
	Shutdown(ctx context.Context) error
}
