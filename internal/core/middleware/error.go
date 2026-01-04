package middleware

import (
	"context"

	"github.com/Pencil-Boot/go-http-server-sdk/internal/core/router"
	httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"
)

type errorHandler struct {
}

func NewErrorHandler() httpservercontract.Middleware {
	return &errorHandler{}
}

func (e *errorHandler) Handle(ctx context.Context, req httpservercontract.Request, next httpservercontract.Handler) (resp httpservercontract.Response, err error) {
	defer func() {
		if r := recover(); r != nil {
			resp = router.NewResponse(500, errorResponse{
				Code:    httpservercontract.HttpServerUnexpectedError,
				Message: "Unexpected error",
			})
			err = nil
		}
	}()

	resp, err = next(ctx, req)
	if err == nil {
		return resp, nil
	}

	if httpErr, ok := err.(httpservercontract.HTTPError); ok {
		return router.NewResponse(httpErr.Status(), errorResponse{
			Code:    httpErr.Code(),
			Message: httpErr.Message(),
			Causes:  httpErr.Causes(),
		}), nil
	}

	//TODO: should use status from core lib
	return router.NewResponse(500, errorResponse{
		Code:    httpservercontract.HttpServerUnexpectedError,
		Message: "Unexpected error",
	}), nil
}

type errorResponse struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Causes  []string `json:"causes,omitempty"`
}
