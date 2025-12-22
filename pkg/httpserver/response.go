package httpserver

import (
	"go-http-server-sdk/internal/core/router"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"
)

func NewResponse(status int, data any) httpservercontract.Response {
	return router.NewResponse(status, data)
}
