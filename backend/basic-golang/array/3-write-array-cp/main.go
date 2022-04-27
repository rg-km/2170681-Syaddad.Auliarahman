package main

import "fmt"

// Kalian udah tau tuh cara memasukan nilai ke array.
// Coba kalian masukan nama kalian ke array.
// inisialisasi array nya menggunakan "var" ya.
// Lalu outputkan array nya.
func main() {
	// TODO: answer here
	var nama []string
	for i := 0; i < 7; i++ {
		var x string
		fmt.Scanln(&x)
		nama = append(nama, x)
	}
	fmt.Println(nama)
}
