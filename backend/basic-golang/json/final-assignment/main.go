package main

import (
	"encoding/json"
	"log"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here
type Ruang struct {
	RuangTamu Items `json:"ruangTamu"`
}

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	Items []Meja `json:"items"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func NewRuang(r Ruang) Ruang {
	return r
}
