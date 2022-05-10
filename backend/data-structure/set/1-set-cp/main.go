// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import (
	"fmt"
)

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "Java", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
	fmt.Println(RemoveDuplicates(str1))
}

func Intersection(str1, str2 []string) (inter []string) {
	temp := make(map[string]bool)
	var res []string
	for _, val2 := range str2 {
		if _, ok := temp[val2]; !ok {
			temp[val2] = true
		}
	}
	for _, val1 := range str1 {
		if _, ok := temp[val1]; ok {
			res = append(res, val1)
		}
	}

	return res // TODO: replace this
}

func RemoveDuplicates(elements []string) (nodups []string) {
	// var ex int
	// ex = 0
	// for i := 0; i < len(elements); i++ {
	// 	for j := 0; j < len(elements); j++ {
	// 		//fmt.Println(string(elements[i]), string(elements[j]))
	// 		if elements[i] == elements[j] {
	// 			ex += 1
	// 			fmt.Println(string(elements[i]), string(elements[j]))
	// 		}
	// 	}
	// }
	temp := make(map[string]bool)
	var res []string

	// for _, item := range elements {
	// 	temp[item] = true
	// }

	for _, x := range elements {
		if _, exist := temp[x]; !exist {
			temp[x] = true
			res = append(res, x)
		}
	}

	return res // TODO: replace this
}
