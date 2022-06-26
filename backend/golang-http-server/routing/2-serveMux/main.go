package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(AuthErrorResponse{Error: "No token provided"})
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
