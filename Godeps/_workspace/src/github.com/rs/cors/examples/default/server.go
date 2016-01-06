package main

import (
	"net/http"

	"asana-task-bot/Godeps/_workspace/src/github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Use default options
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
