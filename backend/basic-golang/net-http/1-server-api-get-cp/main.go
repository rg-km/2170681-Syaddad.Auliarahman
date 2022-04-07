package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Dari contoh sebelumnya tambahkan filter untuk
// menampilkan meja berdasarkan total meja
// key input dari client menggunakan `total`
// contuh URL: http://localhost:8080/table?total=2

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		// TODO: answer here
		total := r.FormValue("total")
		intTotal, _ := strconv.Atoi(total)
		var resultAll []Table
		if total == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid total\n"))
			return
		}
		for _, table := range data {
			if table.Total == intTotal {
				resultAll = append(resultAll, table)
			}
		}
		result, err := json.Marshal(resultAll)
		if err != nil {
			http.Error(w, "invalid total\n", http.StatusInternalServerError)
			return
		}
		if len(resultAll) > 0 {
			w.WriteHeader(http.StatusOK)
			w.Write(result)
			return
		}

		http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
		//http.Error(w, "", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/table", TableHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
