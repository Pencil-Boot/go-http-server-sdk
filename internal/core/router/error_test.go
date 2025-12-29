package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("should create a valid HTTPError", func(t *testing.T) {
		status := 400
		code := "ERR001"
		message := "Bad Request"
		causes := []string{"invalid field"}

		err := NewError(status, code, message, causes)

		assert.NotNil(t, err)
		assert.Equal(t, status, err.Status())
		assert.Equal(t, code, err.Code())
		assert.Equal(t, message, err.Message())
		assert.Equal(t, causes, err.Causes())
	})

	t.Run("should create a valid HTTPError with empty causes", func(t *testing.T) {
		status := 500
		code := "ERR002"
		message := "Internal Error"
		causes := []string{}

		err := NewError(status, code, message, causes)

		assert.NotNil(t, err)
		assert.Equal(t, status, err.Status())
		assert.Equal(t, code, err.Code())
		assert.Equal(t, message, err.Message())
		assert.Equal(t, causes, err.Causes())
	})

	t.Run("should panic if causes is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			NewError(500, "ERR003", "Panic", nil)
		})
	})
}
