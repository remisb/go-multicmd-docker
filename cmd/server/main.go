package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHander)
	http.ListenAndServe(":8081", mux)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server index page"))
}
