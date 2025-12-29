package echo

import (
	"encoding/json"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"io"

	"github.com/labstack/echo/v4"
)

// echoRequest implements httpserver.Request using Echo context
type echoRequest struct {
	ctx         echo.Context
	headers     httpservercontract.RequestHeaders
	params      httpservercontract.Params
	queryParams httpservercontract.QueryParams
}

// newEchoRequest creates a new request from Echo context
func newEchoRequest(c echo.Context) httpservercontract.Request {
	return &echoRequest{
		ctx: c,
	}
}

// Method returns the HTTP method
func (r *echoRequest) Method() string {
	return r.ctx.Request().Method
}

// Path returns the request path
func (r *echoRequest) Path() string {
	return r.ctx.Request().URL.Path
}

// Headers returns the request headers
func (r *echoRequest) Headers() httpservercontract.RequestHeaders {
	if r.headers == nil {
		r.headers = &echoRequestHeaders{ctx: r.ctx}
	}
	return r.headers
}

// Header returns a single header value
func (r *echoRequest) Header(key string) string {
	return r.ctx.Request().Header.Get(key)
}

// Params returns the path parameters
func (r *echoRequest) Params() httpservercontract.Params {
	if r.params == nil {
		r.params = &echoParams{ctx: r.ctx}
	}
	return r.params
}

// Param returns a single path parameter
func (r *echoRequest) Param(key string) string {
	return r.ctx.Param(key)
}

// QueryParams returns the query parameters
func (r *echoRequest) QueryParams() httpservercontract.QueryParams {
	if r.queryParams == nil {
		r.queryParams = &echoQueryParams{ctx: r.ctx}
	}
	return r.queryParams
}

// QueryParam returns a single query parameter
func (r *echoRequest) QueryParam(key string) string {
	return r.ctx.QueryParam(key)
}

// Bind binds the request body to the provided pointer
func (r *echoRequest) Bind(ptr any) error {
	// Read body
	body, err := io.ReadAll(r.ctx.Request().Body)
	if err != nil {
		return err
	}

	// Try to bind as JSON
	return json.Unmarshal(body, ptr)
}

// echoRequestHeaders implements httpserver.RequestHeaders
type echoRequestHeaders struct {
	ctx echo.Context
}

func (h *echoRequestHeaders) Get(key string) []string {
	return h.ctx.Request().Header.Values(key)
}

func (h *echoRequestHeaders) All() map[string][]string {
	return h.ctx.Request().Header
}

// echoParams implements httpserver.Params
type echoParams struct {
	ctx echo.Context
}

func (p *echoParams) Get(key string) string {
	return p.ctx.Param(key)
}

func (p *echoParams) All() map[string]string {
	params := make(map[string]string)
	for _, param := range p.ctx.ParamNames() {
		params[param] = p.ctx.Param(param)
	}
	return params
}

// echoQueryParams implements httpserver.QueryParams
type echoQueryParams struct {
	ctx echo.Context
}

func (q *echoQueryParams) Get(key string) []string {
	return q.ctx.QueryParams()[key]
}

func (q *echoQueryParams) All() map[string][]string {
	return q.ctx.QueryParams()
}
