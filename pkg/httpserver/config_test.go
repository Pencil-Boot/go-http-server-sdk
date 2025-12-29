package httpserver

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigWrappers(t *testing.T) {
	t.Run("WithReadTimeout", func(t *testing.T) {
		opt := WithReadTimeout(10 * time.Second)
		assert.NotNil(t, opt)
	})

	t.Run("WithWriteTimeout", func(t *testing.T) {
		opt := WithWriteTimeout(10 * time.Second)
		assert.NotNil(t, opt)
	})

	t.Run("WithIdleTimeout", func(t *testing.T) {
		opt := WithIdleTimeout(10 * time.Second)
		assert.NotNil(t, opt)
	})

	t.Run("WithMaxHeaderBytes", func(t *testing.T) {
		opt := WithMaxHeaderBytes(1024)
		assert.NotNil(t, opt)
	})

	t.Run("WithTimeout", func(t *testing.T) {
		opt := WithTimeout(10 * time.Second)
		assert.NotNil(t, opt)
	})

	t.Run("WithMiddlewares", func(t *testing.T) {
		opt := WithMiddlewares(nil)
		assert.NotNil(t, opt)
	})
}
