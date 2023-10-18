package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	filepathRoot := "."
	apiCfg := &apiConfig{
		fileserverHits: 0,
	}

	r := chi.NewRouter()
	// Define tus rutas aquí
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	r.Handle("/app", fsHandler)
	r.Handle("/app/*", fsHandler)
	r.HandleFunc("/hola", holaMundo)
	//r.Get("/api/healthz", healthCheckHandler)
	//r.Get("/api/metrics", apiCfg.handlerMetrics)
	//r.Get("/reset", apiCfg.resetHandler)

	// Create a new router to bind and register the /healthz, /reset and /metrics endpoints on, and then r.Mount() that router at /api in our main router.
	apiRouter := func() http.Handler {
		r := chi.NewRouter()
		r.Get("/healthz", healthCheckHandler)
		r.Get("/metrics", apiCfg.handlerMetrics)
		r.Get("/reset", apiCfg.resetHandler)
		return r
	}
	r.Mount("/api", apiRouter())

	// Ahora envuelve mux con el middleware CORS
	corsMux := middlewareCors(r)

	http.ListenAndServe(":8080", corsMux)
}
