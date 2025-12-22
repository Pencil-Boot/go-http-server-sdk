package httpservercontract

import (
	"time"
)

type Config interface {
	// ReadTimeout returns the read timeout for this server
	ReadTimeout() time.Duration

	// WriteTimeout returns the write timeout for this server
	WriteTimeout() time.Duration

	// IdleTimeout returns the idle timeout for this server
	IdleTimeout() time.Duration

	// MaxHeaderBytes returns the max header bytes for this server
	MaxHeaderBytes() int
}

type config struct {
	readTimeout    time.Duration
	writeTimeout   time.Duration
	idleTimeout    time.Duration
	maxHeaderBytes int
}

type Option func(*config)

func NewConfig(opts ...Option) Config {
	c := &config{
		readTimeout:    15 * time.Second,
		writeTimeout:   15 * time.Second,
		idleTimeout:    60 * time.Second,
		maxHeaderBytes: 1 << 20, // 1 MB
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *config) ReadTimeout() time.Duration {
	return c.readTimeout
}

func (c *config) WriteTimeout() time.Duration {
	return c.writeTimeout
}

func (c *config) IdleTimeout() time.Duration {
	return c.idleTimeout
}

func (c *config) MaxHeaderBytes() int {
	return c.maxHeaderBytes
}

func WithReadTimeout(d time.Duration) Option {
	return func(b *config) {
		b.readTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(b *config) {
		b.writeTimeout = d
	}
}

func WithIdleTimeout(d time.Duration) Option {
	return func(b *config) {
		b.idleTimeout = d
	}
}

func WithMaxHeaderBytes(n int) Option {
	return func(b *config) {
		b.maxHeaderBytes = n
	}
}
