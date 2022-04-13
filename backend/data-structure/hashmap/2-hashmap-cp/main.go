// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {
	x := map[rune]bool{}
	anagram := "Anagram"
	Banagram := "Bukan Anagram"
	if len(str1) != len(str2) {
		return Banagram
	}
	for _, val := range str1 {
		x[val] = true
	}
	//lenStr1 := len(str1)
	//akhirStr1 := str1[lenStr1+1]
	for _, val := range str2 {
		//fmt.Printf("%c", val)
		if _, exist := x[val]; exist {
			if x[val] != false {
				x[val] = false
			}
		}
	}
	for _, val := range x {
		if val == true {
			return Banagram
		}
	}

	return anagram // TODO: replace this
}
