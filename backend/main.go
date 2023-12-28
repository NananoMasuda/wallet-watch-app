// backend/main.go
package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Your API routes here

	handler := cors.Default().Handler(router)

	http.Handle("/", handler)
	http.ListenAndServe(":8000", nil)
}