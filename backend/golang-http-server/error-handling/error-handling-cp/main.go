package main

import (
	"log"
	"net/http"
)

// Dari contoh yang telah diberikan, lengkapi handler dibawah ini.
// Lakukan pengecekan terhadap request method, jika request method tidak sama dengan GET maka return method not allowed.
// Selain itu gunakan variabel data dan err, jika err terdapat error, maka handler akan return internal server error.
// Jika panjang data dari data sama dengan 0, maka return not found.

var data []string
var err error

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	for _, v := range data {
		w.Write([]byte(v + "\n"))
	}

}

func main() {

	// Menambahkan handler pada objek server
	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
