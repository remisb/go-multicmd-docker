package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHander)
	http.ListenAndServe(":8080", mux)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("admin index page"))
}
