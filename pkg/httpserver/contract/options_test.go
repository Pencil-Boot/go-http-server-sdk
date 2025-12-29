package httpservercontract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OptionsSuite struct {
	suite.Suite
}

func TestOptionsSuite(t *testing.T) {
	suite.Run(t, new(OptionsSuite))
}

func (s *OptionsSuite) TestServerOptions() {
	tests := []struct {
		name     string
		option   ServerOption
		validate func(*ServerConfig)
	}{
		{
			name:   "PortOption",
			option: NewPortOption(9090),
			validate: func(c *ServerConfig) {
				assert.Equal(s.T(), 9090, c.port)
			},
		},
		{
			name:   "ReadTimeoutOption",
			option: NewReadTimeoutOption(10 * time.Second),
			validate: func(c *ServerConfig) {
				assert.Equal(s.T(), 10*time.Second, c.readTimeout)
			},
		},
		{
			name:   "WriteTimeoutOption",
			option: NewWriteTimeoutOption(20 * time.Second),
			validate: func(c *ServerConfig) {
				assert.Equal(s.T(), 20*time.Second, c.writeTimeout)
			},
		},
		{
			name:   "IdleTimeoutOption",
			option: NewIdleTimeoutOption(30 * time.Second),
			validate: func(c *ServerConfig) {
				assert.Equal(s.T(), 30*time.Second, c.idleTimeout)
			},
		},
		{
			name:   "MaxHeaderBytesOption",
			option: NewMaxHeaderBytesOption(2048),
			validate: func(c *ServerConfig) {
				assert.Equal(s.T(), 2048, c.maxHeaderBytes)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			cfg := &ServerConfig{}
			tt.option.applyServer(cfg)
			tt.validate(cfg)
		})
	}
}

func (s *OptionsSuite) TestRouterOptions() {
	tests := []struct {
		name     string
		option   RouterOption
		validate func(*RouterConfig)
	}{
		{
			name:   "TimeoutOption",
			option: NewTimeoutOption(5 * time.Second),
			validate: func(c *RouterConfig) {
				assert.Equal(s.T(), 5*time.Second, c.timeout)
			},
		},
		{
			name:   "MiddlewaresOption",
			option: NewMiddlewaresOption(nil, nil),
			validate: func(c *RouterConfig) {
				assert.Len(s.T(), c.middlewares, 2)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			cfg := &RouterConfig{}
			tt.option.applyRouter(cfg)
			tt.validate(cfg)
		})
	}
}
