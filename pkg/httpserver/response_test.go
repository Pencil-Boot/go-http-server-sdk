package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	resp := NewResponse(200, "ok")
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.Status())
	assert.Equal(t, "ok", resp.Body())
}
