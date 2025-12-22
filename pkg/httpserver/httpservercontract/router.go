package httpservercontract

type Router interface {
	GET(path string, handler Handler, opts ...Option)
	POST(path string, handler Handler, opts ...Option)
	PUT(path string, handler Handler, opts ...Option)
	DELETE(path string, handler Handler, opts ...Option)
	PATCH(path string, handler Handler, opts ...Option)
	OPTIONS(path string, handler Handler, opts ...Option)
	HEAD(path string, handler Handler, opts ...Option)

	Group(prefix string) Router
	Use(middleware ...Middleware)
}
