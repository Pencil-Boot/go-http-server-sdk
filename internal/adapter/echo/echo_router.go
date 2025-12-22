package echo

import (
	"context"
	"go-http-server-sdk/internal/core/router"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"

	"github.com/labstack/echo/v4"
)

// echoRouter implements contracts.Router using Echo
type echoRouter struct {
	echo *echo.Echo
	// For grouped routers
	group *echo.Group
}

// NewEchoRouter creates a new router using Echo
func NewEchoRouter() httpservercontract.Router {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	return &echoRouter{
		echo: e,
	}
}

// newGroupRouter creates a router from an Echo group
func newGroupRouter(group *echo.Group) httpservercontract.Router {
	return &echoRouter{
		group: group,
	}
}

// GET registers a handler for GET requests
func (r *echoRouter) GET(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("GET", path, handler, opts...)
}

// POST registers a handler for POST requests
func (r *echoRouter) POST(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("POST", path, handler, opts...)
}

// PUT registers a handler for PUT requests
func (r *echoRouter) PUT(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("PUT", path, handler, opts...)
}

// DELETE registers a handler for DELETE requests
func (r *echoRouter) DELETE(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("DELETE", path, handler, opts...)
}

// PATCH registers a handler for PATCH requests
func (r *echoRouter) PATCH(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("PATCH", path, handler, opts...)
}

// OPTIONS registers a handler for OPTIONS requests
func (r *echoRouter) OPTIONS(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("OPTIONS", path, handler, opts...)
}

// HEAD registers a handler for HEAD requests
func (r *echoRouter) HEAD(path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	r.register("HEAD", path, handler, opts...)
}

// Group creates a new router group with the given prefix
func (r *echoRouter) Group(prefix string) httpservercontract.Router {
	if r.group != nil {
		return newGroupRouter(r.group.Group(prefix))
	}
	return newGroupRouter(r.echo.Group(prefix))
}

// Use registers middleware(s)
func (r *echoRouter) Use(middleware ...httpservercontract.Middleware) {
	for _, mw := range middleware {
		echoMw := convertMiddleware(mw)
		if r.group != nil {
			r.group.Use(echoMw)
		} else {
			r.echo.Use(echoMw)
		}
	}
}

// register is a helper method to register routes
func (r *echoRouter) register(method, path string, handler httpservercontract.Handler, opts ...httpservercontract.Option) {
	echoHandler := convertHandler(handler)

	if r.group != nil {
		r.group.Add(method, path, echoHandler)
	} else {
		r.echo.Add(method, path, echoHandler)
	}
}

// GetEcho returns the underlying Echo instance (used by server)
func (r *echoRouter) GetEcho() *echo.Echo {
	return r.echo
}

// convertHandler converts a httpservercontract.Handler to echo.HandlerFunc
func convertHandler(h httpservercontract.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := newEchoRequest(c)
		resp := h(ctx, req)
		return writeResponse(c, resp)
	}
}

// convertMiddleware converts a httpservercontract.Middleware to echo.MiddlewareFunc
func convertMiddleware(mw httpservercontract.Middleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			req := newEchoRequest(c)

			// Create a handler that calls the next Echo handler
			nextHandler := func(ctx context.Context, req httpservercontract.Request) httpservercontract.Response {
				// Execute the next handler
				err := next(c)
				if err != nil {
					// Convert error to response
					return router.NewResponse(500, map[string]string{"error": err.Error()})
				}
				// Return empty response (Echo already wrote it)
				return router.NewResponse(200, nil)
			}

			// Call the middleware
			resp := mw.Handle(ctx, req, nextHandler)

			// Write response if it's not empty
			if resp != nil && resp.Body() != nil {
				return writeResponse(c, resp)
			}

			return nil
		}
	}
}
