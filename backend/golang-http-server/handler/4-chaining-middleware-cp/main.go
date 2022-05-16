package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan chaining middleware.
// Lengkapi middleware AdminMiddleware dan RequestMethodGetMiddleware.
// Pada AdminMiddleware, lakukan pengecekan terhadap header "role". Jika nilai header "role" tidak sama dengan "ADMIN", maka return error HTTP Unauthorized.
// Pada RequestMethodGetMiddleware, lakukan pengecekan terhadap request method. Jika request method selain GET, maka return error HTTP Status Not Allowed.

func TimeHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		output := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		fmt.Fprint(writer, output)
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Lengkapi middleware AdminMiddleware dan RequestMethodGetMiddleware.
		// Pada RequestMethodGetMiddleware, lakukan pengecekan terhadap request method. Jika request method selain GET, maka return error HTTP Status Not Allowed.

		if r.Method != "GET" {
			http.Error(w, "Method notallowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
	// TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pada AdminMiddleware, lakukan pengecekan terhadap header "role". Jika nilai header "role" tidak sama dengan "ADMIN", maka return error HTTP Unauthorized.
		if r.Method != "Role" {
			http.Error(w, "Method notallowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
	// TODO: replace this
}

func main() {
	handler := TimeHandler{
		Timezone: "Asia/Jakarta",
		Format:   time.RFC3339,
	}

	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: getRequestMiddleware(handler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TODO: answer here
