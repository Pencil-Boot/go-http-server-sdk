package echo

import (
	"regexp"

	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/router"
	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"

	"github.com/labstack/echo/v4"
	echoLib "github.com/labstack/echo/v4"
)

var pathParamRegexp = regexp.MustCompile(`\{([^}]+)\}`)

type echoAdapter interface {
	Add(method, path string, handler echoLib.HandlerFunc, middleware ...echoLib.MiddlewareFunc) *echoLib.Route
	Group(prefix string, m ...echoLib.MiddlewareFunc) *echoLib.Group
}

// echoRouter implements router.RouterRegister using Echo
type echoRouter struct {
	echo echoAdapter
}

func NewEchoRouter(echo echoAdapter) router.RouterRegister {
	return &echoRouter{
		echo: echo,
	}
}

// Group creates a new router group with the given prefix
func (r *echoRouter) Group(prefix string) router.RouterRegister {
	normalizedPrefix := pathParamRegexp.ReplaceAllString(prefix, ":$1")
	group := r.echo.Group(normalizedPrefix)
	return NewEchoRouter(group)
}

func (r *echoRouter) Register(
	method httpservercontract.HttpMethod,
	path string,
	handler httpservercontract.Handler,
	config *httpservercontract.RouterConfig,
) {
	echoHandler := convertHandler(handler)
	normalizedPath := pathParamRegexp.ReplaceAllString(path, ":$1")
	r.echo.Add(string(method), normalizedPath, echoHandler)
}

// convertHandler converts a httpservercontract.Handler to echo.HandlerFunc
func convertHandler(h httpservercontract.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := newEchoRequest(c)
		resp, err := h(ctx, req)
		// This error should never happen. If it does, it's a bug in the core layer
		if err != nil {
			//TODO: log error. Should not happen. User status da lib core
			return echo.NewHTTPError(500, "internal error")
		}
		return writeResponse(c, resp)
	}
}
