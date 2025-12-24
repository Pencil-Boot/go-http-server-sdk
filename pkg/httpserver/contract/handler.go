package httpservercontract

import "context"

type Handler func(ctx context.Context, req Request) (Response, error)

// TODO: hold add order method? To allow user to define the order of middleware execution
type Middleware interface {
	Handle(ctx context.Context, req Request, next Handler) (Response, error)
}
