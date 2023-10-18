package main

import (
	"fmt"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

//Create a new handler that writes the number of requests that have been counted as plain text in this format to the HTTP response:

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	//cfg.fileserverHits++
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits)))
}

// Create a new handler that resets the counter to zero and returns a 204 No Content response.
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	cfg.fileserverHits = 0

}
