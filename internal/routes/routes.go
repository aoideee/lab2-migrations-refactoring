package routes

import (
	"net/http"
	"github.com/aoideee/lab2-tyshadaniels/internal/handlers"
	"github.com/aoideee/lab2-tyshadaniels/internal/middleware"
)

func SetupRoutes(mux *http.ServeMux) {
	// Wrap the handlers with the middleware
	// This ensures every request goes through the Logger AND the Counter
	
	// Helper to chain middleware: Logger -> Counter -> Handler
	chain := func(h http.HandlerFunc) http.Handler {
		return middleware.LoggingMiddleware(
			middleware.RequestCountMiddleware(http.HandlerFunc(h)),
		)
	}

	// Registers handlers using the chain
	mux.Handle("/", chain(handlers.Home))
	mux.Handle("/about", chain(handlers.About))
	mux.Handle("/contact", chain(handlers.Contact))
	mux.Handle("/quote", chain(handlers.Quote))
}