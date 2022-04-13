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

// import "fmt"

// func main() {
// 	var str1 = "fried"
// 	var str2 = "fired"
// 	fmt.Println(AnagramsChecker(str1, str2))
// }

func AnagramsChecker(str1 string, str2 string) string {
	var output string
	if len(str1) != len(str2) {
		output = "Bukan Anagram"
	} else {
		var str1Map = map[rune]int{}
		var str2Map = map[rune]int{}

		for _, val := range str1 {
			str1Map[val]++
		}

		for _, val := range str2 {
			str2Map[val]++
		}
		for key, val := range str1Map {
			if str2Map[key] != val {
				output = "Bukan Anagram"
				break
			}
		}
		if output == "" {
			output = "Anagram"
		}
	}
	return output // TODO: replace this
}
