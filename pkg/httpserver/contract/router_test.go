package httpservercontract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRouterConfig(t *testing.T) {
	cfg := NewRouterConfig()
	assert.Equal(t, 15*time.Second, cfg.Timeout())
	assert.Empty(t, cfg.Middlewares())

	cfg2 := NewRouterConfig(NewTimeoutOption(5 * time.Second))
	assert.Equal(t, 5*time.Second, cfg2.Timeout())
}
