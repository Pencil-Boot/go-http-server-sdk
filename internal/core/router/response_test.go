package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPResponse(t *testing.T) {
	resp := NewResponse(201, "created")

	assert.Equal(t, 201, resp.Status())
	assert.Equal(t, "created", resp.Body())
	assert.NotNil(t, resp.Headers())
}

func TestResponseHeaders(t *testing.T) {
	h := NewResponseHeaders()

	h.Set("X-Test", "1")
	assert.Equal(t, []string{"1"}, h.Get("X-Test"))

	h.Add("X-Test", "2")
	assert.Equal(t, []string{"1", "2"}, h.Get("X-Test"))

	assert.Equal(t, map[string][]string{"X-Test": {"1", "2"}}, h.All())
}
