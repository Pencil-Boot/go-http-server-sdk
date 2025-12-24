package httpservercontract

import "time"

type portOption struct {
	value int
}

func NewPortOption(value int) portOption {
	return portOption{value: value}
}

func (o portOption) applyServer(c *ServerConfig) {
	c.port = o.value
}

type readTimeoutOption struct {
	value time.Duration
}

func NewReadTimeoutOption(value time.Duration) readTimeoutOption {
	return readTimeoutOption{value: value}
}

func (o readTimeoutOption) applyServer(c *ServerConfig) {
	c.readTimeout = o.value
}

type writeTimeoutOption struct {
	value time.Duration
}

func NewWriteTimeoutOption(value time.Duration) writeTimeoutOption {
	return writeTimeoutOption{value: value}
}

func (o writeTimeoutOption) applyServer(c *ServerConfig) {
	c.writeTimeout = o.value
}

type idleTimeoutOption struct {
	value time.Duration
}

func NewIdleTimeoutOption(value time.Duration) idleTimeoutOption {
	return idleTimeoutOption{value: value}
}

func (o idleTimeoutOption) applyServer(c *ServerConfig) {
	c.idleTimeout = o.value
}

type maxHeaderBytesOption struct {
	value int
}

func NewMaxHeaderBytesOption(value int) maxHeaderBytesOption {
	return maxHeaderBytesOption{value: value}
}

func (o maxHeaderBytesOption) applyServer(c *ServerConfig) {
	c.maxHeaderBytes = o.value
}

type timeoutOption struct {
	value time.Duration
}

func NewTimeoutOption(value time.Duration) timeoutOption {
	return timeoutOption{value: value}
}

func (o timeoutOption) applyRouter(c *RouterConfig) {
	c.timeout = o.value
}

type middlewareOption struct {
	value []Middleware
}

func NewMiddlewaresOption(value ...Middleware) middlewareOption {
	return middlewareOption{value: value}
}

func (o middlewareOption) applyRouter(c *RouterConfig) {
	c.middlewares = append(c.middlewares, o.value...)
}
