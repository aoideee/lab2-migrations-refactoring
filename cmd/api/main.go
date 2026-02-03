// Filename: cmd/api/main.go

package main

import (
	"log"
	"net/http"
	"github.com/aoideee/lab2-tyshadaniels/internal/routes"
)

func main() {
	// Create a new ServeMux (router) to handle HTTP requests
	mux := http.NewServeMux()

	// Setup the application routes using the routes package
	routes.SetupRoutes(mux)

	// Start the server on port 4000 and log any errors
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}