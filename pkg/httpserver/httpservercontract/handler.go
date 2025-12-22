package httpservercontract

import "context"

type Handler func(ctx context.Context, req Request) Response

type Middleware interface {
	Handle(ctx context.Context, req Request, next Handler) Response
}
