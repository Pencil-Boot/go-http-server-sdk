package httpservercontract

import (
	"slices"
	"time"
)

// TODO: Migrar para a lib core
type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	PATCH   HttpMethod = "PATCH"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
)

type Router interface {
	GET(path string, handler Handler, opts ...RouterOption)
	POST(path string, handler Handler, opts ...RouterOption)
	PUT(path string, handler Handler, opts ...RouterOption)
	DELETE(path string, handler Handler, opts ...RouterOption)
	PATCH(path string, handler Handler, opts ...RouterOption)
	OPTIONS(path string, handler Handler, opts ...RouterOption)
	HEAD(path string, handler Handler, opts ...RouterOption)

	Group(prefix string) Router
	Use(middleware ...Middleware)
}

type RouterOption interface {
	applyRouter(*RouterConfig)
}

type RouterConfig struct {
	timeout     time.Duration
	middlewares []Middleware
}

func (c *RouterConfig) Timeout() time.Duration {
	return c.timeout
}

func (c *RouterConfig) Middlewares() []Middleware {
	return slices.Clone(c.middlewares)
}

func NewRouterConfig(opts ...RouterOption) *RouterConfig {
	c := &RouterConfig{
		timeout:     15 * time.Second,
		middlewares: []Middleware{},
	}

	for _, opt := range opts {
		opt.applyRouter(c)
	}

	return c
}
