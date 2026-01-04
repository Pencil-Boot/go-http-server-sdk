package echo

import (
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/router"
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/server"

	"github.com/labstack/echo/v4"
)

type Echo struct {
	router router.RouterRegister
	server server.ServerExecutor
}

func Initialize() *Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	return &Echo{
		router: NewEchoRouter(e),
		server: NewServer(e),
	}
}

func (e *Echo) Router() router.RouterRegister {
	return e.router
}

func (e *Echo) Server() server.ServerExecutor {
	return e.server
}
