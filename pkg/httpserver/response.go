package httpserver

import (
	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/router"
	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

func NewResponse(status int, data any) httpservercontract.Response {
	return router.NewResponse(status, data)
}
