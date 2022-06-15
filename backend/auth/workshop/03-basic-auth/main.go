package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json
		err := json.NewDecoder(r.Body).Decode(&creds)

		// TODO: answer here

		// Task: Ambil password dari username yang dipakai untuk login

		// TODO: answer here
		password := creds.Password

		// Task: return unauthorized jika password salah
		if users[creds.Username] != password {
			w.WriteHeader(http.StatusUnauthorized)
		}
		// TODO: answer here

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token

		token, err := bcrypt.GenerateFromPassword([]byte(creds.Username), 10)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// TODO: answer here

		// Task: Set token string kedalam cookie response
		http.SetCookie(w, &http.Cookie{
			Name:  cookieName,
			Value: string(token),
		})
		// TODO: answer here
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
