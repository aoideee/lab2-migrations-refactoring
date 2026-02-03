package middleware

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// requestCounter tracks the total number of requests processed
var requestCounter int64

// LoggingMiddleware logs the HTTP method, path, and timestamp for each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s | Path: %s | Timestamp: %v", r.Method, r.URL.Path, time.Now())
		next.ServeHTTP(w, r)
	})
}

// RequestCountMiddleware tracks the total number of requests
func RequestCountMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Increment the counter
        newCount := atomic.AddInt64(&requestCounter, 1)

        // Log the new count
        log.Printf("Request Count: %d", newCount)
        
		// Call the next handler
        next.ServeHTTP(w, r)
    })
}