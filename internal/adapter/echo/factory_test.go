package echo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	e := Initialize()
	assert.NotNil(t, e)
	assert.NotNil(t, e.Router())
	assert.NotNil(t, e.Server())
}
