package bootstrap

import (
	"go-http-server-sdk/internal/adapter/echo"
	"go-http-server-sdk/internal/core/middleware"
	"go-http-server-sdk/internal/core/router"
	"go-http-server-sdk/internal/core/server"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

func CreateServer(cfg *httpservercontract.ServerConfig) httpservercontract.Server {
	echoAdapter := echo.Initialize()
	//Todo: add telemetry middleware
	r := router.NewRouter(echoAdapter.Router(), middleware.NewErrorHandler())
	return server.NewServer(r, cfg, echoAdapter.Server())
}
