package main

import (
	"fmt"
	"go-multicmd-docker/internal/health"
	"net/http"
)

func main() {
	port := 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/admin", indexHander)
	mux.HandleFunc("/admin/health", health.HealthHandler)
	fmt.Printf("Server started and accessible at localhost:/%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin index page"))
}
