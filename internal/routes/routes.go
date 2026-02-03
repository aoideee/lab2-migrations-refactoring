package routes

import (
	"net/http"
	"github.com/aoideee/lab2-tyshadaniels/internal/handlers"
)

// SetupRoutes configures the application's HTTP routes and maps them to handlers
func SetupRoutes(mux *http.ServeMux) {
	// Register handlers for specific URL patterns
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/contact", handlers.Contact)
	mux.HandleFunc("/quote", handlers.Quote)
}