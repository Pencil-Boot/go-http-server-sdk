package echo

import (
	"context"
	"fmt"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
	"net"
	"net/http"
	"testing"
	"time"

	"os"
	"os/exec"

	echoLib "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Generate certs for testing
	cmd := exec.Command("openssl", "req", "-x509", "-newkey", "rsa:2048", "-keyout", "key.pem", "-out", "cert.pem", "-days", "1", "-nodes", "-subj", "/CN=localhost")
	_ = cmd.Run()

	code := m.Run()

	// Cleanup
	_ = os.Remove("cert.pem")
	_ = os.Remove("key.pem")

	os.Exit(code)
}

func TestFillServerConfig(t *testing.T) {
	e := echoLib.New()
	cfg := httpservercontract.NewServerConfig(
		httpservercontract.NewReadTimeoutOption(10*time.Second),
		httpservercontract.NewWriteTimeoutOption(20*time.Second),
		httpservercontract.NewIdleTimeoutOption(30*time.Second),
		httpservercontract.NewMaxHeaderBytesOption(4096),
	)

	fillServerConfig(cfg, e)

	assert.Equal(t, 10*time.Second, e.Server.ReadTimeout)
	assert.Equal(t, 20*time.Second, e.Server.WriteTimeout)
	assert.Equal(t, 30*time.Second, e.Server.IdleTimeout)
	assert.Equal(t, 4096, e.Server.MaxHeaderBytes)
}

func TestGetAddress(t *testing.T) {
	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(9090))
	addr := getAddress(cfg)
	assert.Equal(t, ":9090", addr)
}

func TestEchoServer_Shutdown(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestEchoServer_Start(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)

	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(0)) // Random port

	// Start in a goroutine
	go func() {
		_ = s.Start(cfg)
	}()

	// Give it a moment to start
	time.Sleep(100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestEchoServer_StartTLS(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)

	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(0))

	// Should fail because certs don't exist
	err := s.StartTLS("nonexistent.crt", "nonexistent.key", cfg)
	assert.Error(t, err)
}

func TestEchoServer_StartError(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)
	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(-1)) // Invalid port
	err := s.Start(cfg)
	assert.Error(t, err)
}

func TestEchoServer_ShutdownError(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)

	// Register a handler that hangs
	e.GET("/hang", func(c echoLib.Context) error {
		time.Sleep(2 * time.Second)
		return c.String(200, "ok")
	})

	// Start server
	l, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)
	port := l.Addr().(*net.TCPAddr).Port
	_ = l.Close()

	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(port))
	go func() {
		_ = s.Start(cfg)
	}()
	time.Sleep(200 * time.Millisecond)

	// Make a request that hangs
	go func() {
		_, _ = http.Get(fmt.Sprintf("http://localhost:%d/hang", port))
	}()
	time.Sleep(100 * time.Millisecond)

	// Shutdown with short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = s.Shutdown(ctx)
	assert.Error(t, err)
}

func TestEchoServer_StartTLS_Closed(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)
	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(0))

	// If we close the echo instance, StartTLS should return ErrServerClosed (or similar)
	// which our code handles as nil error.
	_ = e.Close()

	err := s.StartTLS("cert.pem", "key.pem", cfg)
	assert.NoError(t, err)
}

func TestEchoServer_StartTLS_Success(t *testing.T) {
	e := echoLib.New()
	s := NewServer(e)
	cfg := httpservercontract.NewServerConfig(httpservercontract.NewPortOption(0))

	go func() {
		_ = s.StartTLS("cert.pem", "key.pem", cfg)
	}()
	time.Sleep(100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	assert.NoError(t, err)
}
