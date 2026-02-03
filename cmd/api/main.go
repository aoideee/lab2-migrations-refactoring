// Filename: cmd/api/main.go

package main

import (
	"log"
	"net/http"

	"github.com/aoideee/lab2-tyshadaniels/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}