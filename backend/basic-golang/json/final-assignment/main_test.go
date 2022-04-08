package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	jsonencode "github.com/ruang-guru/playground/backend/basic-golang/json/final-assignment"
)

var _ = Describe("JSON Encode", func() {

	Describe("JSON Encode Array Nested", func() {
		It("encoding string JSON array nested", func() {
			items := jsonencode.Ruang{
				jsonencode.Items{
					[]jsonencode.Meja{
						{
							Nama:   "Meja",
							Warna:  "Coklat",
							Jumlah: 20,
							Ukuran: jsonencode.Ukuran{
								Panjang: "50 cm",
								Tinggi:  "25 cm",
							},
						},
						{
							Nama:   "Kursi",
							Warna:  "Hitam",
							Jumlah: 1,
							Ukuran: jsonencode.Ukuran{
								Panjang: "70 cm",
								Tinggi:  "30 cm",
							},
						},
					},
				}}

			meja := jsonencode.NewRuang(items)
			result := meja.EncodeJSON()
			Expect(result).To(Equal(`{"ruangTamu":{"items":[{"nama":"Meja","jumlah":20,"warna":"Coklat","ukuran":{"panjang":"50 cm","tinggi":"25 cm"}},{"nama":"Kursi","jumlah":1,"warna":"Hitam","ukuran":{"panjang":"70 cm","tinggi":"30 cm"}}]}}`))
		})
	})

})
