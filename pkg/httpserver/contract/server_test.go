package httpservercontract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerConfigSuite struct {
	suite.Suite
}

func TestServerConfigSuite(t *testing.T) {
	suite.Run(t, new(ServerConfigSuite))
}

func (s *ServerConfigSuite) TestNewServerConfig_Defaults() {
	cfg := NewServerConfig()

	assert.Equal(s.T(), 8080, cfg.Port())
	assert.Equal(s.T(), 15*time.Second, cfg.ReadTimeout())
	assert.Equal(s.T(), 15*time.Second, cfg.WriteTimeout())
	assert.Equal(s.T(), 60*time.Second, cfg.IdleTimeout())
	assert.Equal(s.T(), 1<<20, cfg.MaxHeaderBytes())
}

func (s *ServerConfigSuite) TestNewServerConfig_WithOptions() {
	cfg := NewServerConfig(
		NewPortOption(9090),
		NewReadTimeoutOption(5*time.Second),
	)

	assert.Equal(s.T(), 9090, cfg.Port())
	assert.Equal(s.T(), 5*time.Second, cfg.ReadTimeout())
	// Others should remain default
	assert.Equal(s.T(), 15*time.Second, cfg.WriteTimeout())
}

func (s *ServerConfigSuite) TestGetters() {
	cfg := &ServerConfig{
		port:           1234,
		readTimeout:    1 * time.Second,
		writeTimeout:   2 * time.Second,
		idleTimeout:    3 * time.Second,
		maxHeaderBytes: 4096,
	}

	assert.Equal(s.T(), 1234, cfg.Port())
	assert.Equal(s.T(), 1*time.Second, cfg.ReadTimeout())
	assert.Equal(s.T(), 2*time.Second, cfg.WriteTimeout())
	assert.Equal(s.T(), 3*time.Second, cfg.IdleTimeout())
	assert.Equal(s.T(), 4096, cfg.MaxHeaderBytes())
}
