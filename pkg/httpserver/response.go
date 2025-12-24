package httpserver

import (
	"go-http-server-sdk/internal/core/router"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

func NewResponse(status int, data any) httpservercontract.Response {
	return router.NewResponse(status, data)
}
