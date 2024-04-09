package main

import (
	"go-multicmd-docker/internal/health"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHander)
	mux.HandleFunc("/health", health.HealthHandler)
	http.ListenAndServe(":8080", mux)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin index page"))
}
