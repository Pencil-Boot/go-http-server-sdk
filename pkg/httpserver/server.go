package httpserver

import (
	"go-http-server-sdk/internal/bootstrap"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

func NewServer(
	port int,
	opts ...httpservercontract.ServerOption,
) httpservercontract.Server {
	opts = append(opts, httpservercontract.NewPortOption(port))
	cfg := httpservercontract.NewServerConfig(opts...)

	return bootstrap.CreateServer(cfg)
}
