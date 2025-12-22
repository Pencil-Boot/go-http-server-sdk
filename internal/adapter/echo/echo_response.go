package echo

import (
	"encoding/json"
	"go-http-server-sdk/pkg/httpserver/httpservercontract"

	"github.com/labstack/echo/v4"
)

// writeResponse writes a httpserver.Response to the Echo context
func writeResponse(c echo.Context, resp httpservercontract.Response) error {
	if resp == nil {
		return c.NoContent(200)
	}

	// Set status code
	status := resp.Status()
	if status == 0 {
		status = 200
	}

	// Set headers
	if resp.Headers() != nil {
		for key, values := range resp.Headers().All() {
			for _, value := range values {
				c.Response().Header().Add(key, value)
			}
		}
	}

	// Write body
	body := resp.Body()
	if body == nil {
		return c.NoContent(status)
	}

	// Handle different body types
	switch v := body.(type) {
	case string:
		return c.String(status, v)
	case []byte:
		return c.Blob(status, "application/octet-stream", v)
	default:
		// Try to marshal as JSON
		data, err := json.Marshal(body)
		if err != nil {
			return c.String(500, "failed to marshal response")
		}
		return c.JSONBlob(status, data)
	}
}
