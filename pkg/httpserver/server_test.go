package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := NewServer(8080)
	assert.NotNil(t, s)
}
