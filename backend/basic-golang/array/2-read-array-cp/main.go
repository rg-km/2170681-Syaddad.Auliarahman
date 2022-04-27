package main

import "fmt"

// Nah pada checkpoint kali ini kita akan membaca huruf pertama dan terakhir nama kalian
// Yuk dicoba, masih ingat kan dengan cara inisialisasi array pada contoh kode nomor 1 ?
// Outputkan hasilnya ya
func main() {
	// TODO: answer here
	nama := [7]string{"s", "y", "a", "d", "d", "a", "d"}
	fmt.Println(nama[0], nama[len(nama)-1])
}
