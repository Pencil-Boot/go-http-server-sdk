package router

import (
	"context"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

type router struct {
	register    RouterRegister
	middlewares []httpservercontract.Middleware
}

func NewRouter(register RouterRegister, middlewares ...httpservercontract.Middleware) httpservercontract.Router {
	return &router{
		register:    register,
		middlewares: middlewares,
	}
}

func (r *router) GET(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.GET, path, handler, opts...)
}

func (r *router) POST(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.POST, path, handler, opts...)
}

func (r *router) PUT(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.PUT, path, handler, opts...)
}

func (r *router) DELETE(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.DELETE, path, handler, opts...)
}

func (r *router) PATCH(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.PATCH, path, handler, opts...)
}

func (r *router) OPTIONS(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.OPTIONS, path, handler, opts...)
}

func (r *router) HEAD(path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	r.registerHandler(httpservercontract.HEAD, path, handler, opts...)
}

func (r *router) Group(prefix string) httpservercontract.Router {
	return NewRouter(r.register.Group(prefix), r.middlewares...)
}

func (r *router) Use(middleware ...httpservercontract.Middleware) {
	for _, mw := range middleware {
		r.middlewares = append(r.middlewares, mw)
	}
}

func (r *router) registerHandler(method httpservercontract.HttpMethod, path string, handler httpservercontract.Handler, opts ...httpservercontract.RouterOption) {
	config := r.convertToRouterConfig(opts...)
	wrappedHandler := r.wrapHandlerWithMiddlewares(handler, config.Middlewares())
	r.register.Register(method, path, wrappedHandler, config)
}

func (r *router) convertToRouterConfig(opts ...httpservercontract.RouterOption) *httpservercontract.RouterConfig {
	allOpts := make([]httpservercontract.RouterOption, 0, len(opts)+1)
	allOpts = append(allOpts, httpservercontract.NewMiddlewaresOption(r.middlewares...))
	allOpts = append(allOpts, opts...)
	return httpservercontract.NewRouterConfig(allOpts...)
}

func (r *router) wrapHandlerWithMiddlewares(
	handler httpservercontract.Handler,
	middlewares []httpservercontract.Middleware,
) httpservercontract.Handler {
	currentHandler := handler

	for index := len(middlewares) - 1; index >= 0; index-- {
		middleware := middlewares[index]
		next := currentHandler

		currentHandler = func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
			return middleware.Handle(ctx, req, next)
		}
	}

	return currentHandler
}
