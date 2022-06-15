package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Authorization menggunakan claim JWT
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		// TODO: answer here
		expPwd, ok := users[creds.Username]
		if !ok || expPwd.Password != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user invalid password"))
			return
		}
		// Task: 1. Deklarasi expiry time untuk token jwt
		// 		 2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//       3. Expiry time menggunakan time millisecond

		// TODO: answer here
		expTime := time.Now().Add(50000 * time.Millisecond)

		claims := &Claims{
			Username: creds.Username,
			Role:     expPwd.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.Unix(),
			},
		}

		// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

		// TODO: answer here
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//       2. Buat return internal error ketika ada kesalahan ketika pembuatan JWT string

		// TODO: answer here
		tknStr, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		// Task: Set token string kedalam cookie response
		// TODO: answer here
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tknStr,
			Expires: expTime,
		})
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token not found!"))
		}

		// Task: Ambil value dari cookie token

		// TODO: answer here
		tknStr := c.Value

		// Task: Deklarasi variable claim

		// TODO: answer here
		claims := &Claims{}

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		// TODO: answer here
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token not found!"))
			return
		}

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token expired"))
			return
		}

		// Task: return unauthorized ketika role user tidak sesuai dengan role admin

		// TODO: answer here
		if claims.Role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Welcome Admin %s!", claims.Username)))
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token not found!"))
			return
		}
		// Task: Ambil value dari cookie token

		// TODO: answer here
		tknStr := c.Value
		// Task: Deklarasi variable claim

		// TODO: answer here
		claims := &Claims{}
		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		// TODO: answer here
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token not found!"))
			return
		}
		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token expired"))
			return
		}
		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	})

	return mux
}
