package main

import (
	"fmt"
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		param := r.URL.Query().Get("name")

		if param == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		if !IsNameExists(param) {
			http.Error(w, "Name does not exists", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Println(param)
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/name", GetNameHandler())
	// TODO: answer here
	return mux
}
