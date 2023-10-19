package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

//func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
//	//cfg.fileserverHits++
//	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits)))
//}

func (a *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Leer el archivo
	content, err := ioutil.ReadFile("admin.html")
	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusInternalServerError)
		return
	}

	// Reemplaza el marcador de posición con el número de visitas
	response := fmt.Sprintf(string(content), a.fileserverHits)

	w.Write([]byte(response))
}

// Create a new handler that resets the counter to zero and returns a 204 No Content response.
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	cfg.fileserverHits = 0

}

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(params.Body, badWords)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleaned,
	})
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
