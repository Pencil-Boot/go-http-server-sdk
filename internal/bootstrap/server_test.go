package bootstrap

import (
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	cfg := httpservercontract.NewServerConfig()
	s := CreateServer(cfg)

	assert.NotNil(t, s)
	assert.NotNil(t, s.Router())
}
