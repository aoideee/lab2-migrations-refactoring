package routes

import (
	"net/http"

	"github.com/aoideee/lab2-tyshadaniels/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux) {
	
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/contact", handlers.Contact)
	mux.HandleFunc("/quote", handlers.Quote)
}