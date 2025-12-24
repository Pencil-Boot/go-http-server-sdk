package router

import httpservercontract "go-http-server-sdk/pkg/httpserver/contract"

// HTTPResponse is a concrete implementation of the Response interface.
// It represents an HTTP response with status code, headers, and body.
type HTTPResponse struct {
	statusCode int
	headers    httpservercontract.ResponseHeaders
	body       any
}

// Status returns the HTTP status code
func (r *HTTPResponse) Status() int {
	return r.statusCode
}

// Headers returns the response headers
func (r *HTTPResponse) Headers() httpservercontract.ResponseHeaders {
	return r.headers
}

// Body returns the response body
func (r *HTTPResponse) Body() any {
	return r.body
}

// NewResponse creates a new HTTPResponse with the given status code and body
func NewResponse(statusCode int, body any) *HTTPResponse {
	return &HTTPResponse{
		statusCode: statusCode,
		headers:    NewResponseHeaders(),
		body:       body,
	}
}

// responseHeaders is a concrete implementation of ResponseHeaders
type responseHeaders struct {
	headers map[string][]string
}

// NewResponseHeaders creates a new ResponseHeaders instance
func NewResponseHeaders() httpservercontract.ResponseHeaders {
	return &responseHeaders{
		headers: make(map[string][]string),
	}
}

func (h *responseHeaders) Get(key string) []string {
	return h.headers[key]
}

func (h *responseHeaders) Set(key, value string) {
	h.headers[key] = []string{value}
}

func (h *responseHeaders) Add(key, value string) {
	h.headers[key] = append(h.headers[key], value)
}

func (h *responseHeaders) All() map[string][]string {
	return h.headers
}
