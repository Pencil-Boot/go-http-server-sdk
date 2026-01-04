package bootstrap

import (
	"github.com/Pencil-Boot/go-http-server-sdk/internal/adapter/echo"
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/middleware"
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/router"
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/server"
	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

func CreateServer(cfg *httpservercontract.ServerConfig) httpservercontract.Server {
	echoAdapter := echo.Initialize()
	//Todo: add telemetry middleware
	r := router.NewRouter(echoAdapter.Router(), middleware.NewErrorHandler())
	return server.NewServer(r, cfg, echoAdapter.Server())
}
