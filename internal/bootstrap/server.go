package bootstrap

import (
	"go-http-server-sdk/internal/adapter/echo"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"
)

func CreateServer(port int, cfg httpservercontract.Config) httpservercontract.Server {
	return echo.NewServer(port, cfg)
}
