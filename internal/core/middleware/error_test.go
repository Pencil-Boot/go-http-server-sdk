package middleware

import (
	"context"
	"errors"
	"testing"

	"go-http-server-sdk/internal/core/router"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"

	"github.com/stretchr/testify/assert"
)

type mockRequest struct {
	httpservercontract.Request
}

func TestErrorHandler_Handle(t *testing.T) {
	mw := NewErrorHandler()
	ctx := context.Background()
	req := &mockRequest{}

	t.Run("should return response when no error occurs", func(t *testing.T) {
		expectedResp := router.NewResponse(200, "ok")
		next := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
			return expectedResp, nil
		}

		resp, err := mw.Handle(ctx, req, next)

		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("should handle HTTPError and return formatted response", func(t *testing.T) {
		httpErr := router.NewError(400, "ERR001", "Bad Request", []string{"invalid field"})
		next := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
			return nil, httpErr
		}

		resp, err := mw.Handle(ctx, req, next)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 400, resp.Status())

		body, ok := resp.Body().(errorResponse)
		assert.True(t, ok)
		assert.Equal(t, "ERR001", body.Code)
		assert.Equal(t, "Bad Request", body.Message)
		assert.Equal(t, []string{"invalid field"}, body.Causes)
	})

	t.Run("should handle generic error and return 500", func(t *testing.T) {
		next := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
			return nil, errors.New("generic error")
		}

		resp, err := mw.Handle(ctx, req, next)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 500, resp.Status())

		body, ok := resp.Body().(errorResponse)
		assert.True(t, ok)
		assert.Equal(t, httpservercontract.HttpServerUnexpectedError, body.Code)
		assert.Equal(t, "Unexpected error", body.Message)
	})

	t.Run("should recover from panic and return 500", func(t *testing.T) {
		next := func(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
			panic("something went wrong")
		}

		resp, err := mw.Handle(ctx, req, next)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 500, resp.Status())

		body, ok := resp.Body().(errorResponse)
		assert.True(t, ok)
		assert.Equal(t, httpservercontract.HttpServerUnexpectedError, body.Code)
		assert.Equal(t, "Unexpected error", body.Message)
	})
}
