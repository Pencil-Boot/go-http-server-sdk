package httpserver

import (
	"go-http-server-sdk/pkg/httpserver/httpservercontract"
	"time"
)

func WithReadTimeout(timeout time.Duration) httpservercontract.Option {
	return httpservercontract.WithReadTimeout(timeout)
}

func WithWriteTimeout(timeout time.Duration) httpservercontract.Option {
	return httpservercontract.WithWriteTimeout(timeout)
}

func WithIdleTimeout(d time.Duration) httpservercontract.Option {
	return httpservercontract.WithIdleTimeout(d)
}

func WithMaxHeaderBytes(bytes int) httpservercontract.Option {
	return httpservercontract.WithMaxHeaderBytes(bytes)
}
