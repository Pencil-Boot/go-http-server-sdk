package httpserver

import (
	"go-http-server-sdk/internal/bootstrap"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"
)

func NewServer(
	port int,
	opts ...httpservercontract.Option,
) httpservercontract.Server {
	cfg := httpservercontract.NewConfig(opts...)

	return bootstrap.CreateServer(port, cfg)
}
