package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-http-server-sdk/pkg/httpserver"
	httpservercontract "go-http-server-sdk/pkg/httpserver/contract"
)

func main() {
	// Create server
	srv := httpserver.NewServer(
		8080,
		httpserver.WithReadTimeout(15*time.Second),
		httpserver.WithWriteTimeout(15*time.Second),
	)

	// Get router
	router := srv.Router()

	// Register global middleware
	router.Use(&LoggingMiddleware{})

	// Register routes
	router.GET("/hello", helloHandler)
	router.GET("/users/:id", getUserHandler)
	router.POST("/users", createUserHandler)

	// Create API group
	api := router.Group("/api")
	api.GET("/status", statusHandler)

	// v1 group
	v1 := api.Group("/v1")
	v1.GET("/products", productsHandler)
	v1.GET("/products/:id", productDetailHandler)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		log.Println("Server starting on :8080...")
		log.Println("Try:")
		log.Println("  curl http://localhost:8080/hello")
		log.Println("  curl http://localhost:8080/users/123")
		log.Println("  curl -X POST http://localhost:8080/users -d '{\"name\":\"João\",\"email\":\"joao@example.com\"}'")
		log.Println("  curl http://localhost:8080/api/status")
		log.Println("  curl http://localhost:8080/api/v1/products")
		log.Println("")

		if err := srv.Start(); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	log.Println("\nShutting down server...")

	// Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}

// Handlers

func helloHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	return httpserver.NewResponse(200, map[string]string{
		"message": "Hello, World!",
	}), nil
}

func getUserHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	userID := req.Param("id")
	return httpserver.NewResponse(200, map[string]interface{}{
		"user_id": userID,
		"name":    fmt.Sprintf("User %s", userID),
		"email":   fmt.Sprintf("user%s@example.com", userID),
	}), nil
}

func createUserHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	var body map[string]string
	if err := req.Bind(&body); err != nil {
		return httpserver.NewResponse(400, map[string]string{
			"error": "Invalid request body",
		}), err
	}

	return httpserver.NewResponse(201, map[string]interface{}{
		"id":      "new-user-id",
		"name":    body["name"],
		"email":   body["email"],
		"created": true,
	}), nil
}

func statusHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	return httpserver.NewResponse(200, map[string]string{
		"status": "ok",
	}), nil
}

func productsHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	products := []map[string]interface{}{
		{"id": "1", "name": "Product 1", "price": 100},
		{"id": "2", "name": "Product 2", "price": 200},
	}
	return httpserver.NewResponse(200, products), nil
}

func productDetailHandler(ctx context.Context, req httpservercontract.Request) (httpservercontract.Response, error) {
	productID := req.Param("id")
	return httpserver.NewResponse(200, map[string]interface{}{
		"id":    productID,
		"name":  fmt.Sprintf("Product %s", productID),
		"price": 150,
	}), nil
}

// Middleware

type LoggingMiddleware struct{}

func (m *LoggingMiddleware) Handle(ctx context.Context, req httpservercontract.Request, next httpservercontract.Handler) (httpservercontract.Response, error) {
	start := time.Now()

	log.Printf("[%s] %s - Started", req.Method(), req.Path())

	resp, err := next(ctx, req)

	duration := time.Since(start)
	log.Printf("[%s] %s - Completed in %v with status %d",
		req.Method(), req.Path(), duration, resp.Status())

	return resp, err
}
