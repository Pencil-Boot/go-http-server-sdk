package httpserver

import (
	"github.com/Pencil-Boot/go-http-server-sdk/internal/bootstrap"
	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

func NewServer(
	port int,
	opts ...httpservercontract.ServerOption,
) httpservercontract.Server {
	opts = append(opts, httpservercontract.NewPortOption(port))
	cfg := httpservercontract.NewServerConfig(opts...)

	return bootstrap.CreateServer(cfg)
}
