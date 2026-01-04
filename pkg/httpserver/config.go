package httpserver

import (
	"time"

	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

func WithReadTimeout(timeout time.Duration) httpservercontract.ServerOption {
	return httpservercontract.NewReadTimeoutOption(timeout)
}

func WithWriteTimeout(timeout time.Duration) httpservercontract.ServerOption {
	return httpservercontract.NewWriteTimeoutOption(timeout)
}

func WithIdleTimeout(timeout time.Duration) httpservercontract.ServerOption {
	return httpservercontract.NewIdleTimeoutOption(timeout)
}

func WithMaxHeaderBytes(bytes int) httpservercontract.ServerOption {
	return httpservercontract.NewMaxHeaderBytesOption(bytes)
}

func WithTimeout(timeout time.Duration) httpservercontract.RouterOption {
	return httpservercontract.NewTimeoutOption(timeout)
}

func WithMiddlewares(middlewares ...httpservercontract.Middleware) httpservercontract.RouterOption {
	return httpservercontract.NewMiddlewaresOption(middlewares...)
}

//TODO: add option to rate limit
