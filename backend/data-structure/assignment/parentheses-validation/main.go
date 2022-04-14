package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// TODO: answer here
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	parentthesesMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := stack.Stack{
		Top:  -1,
		Data: []rune{},
	}

	for _, v := range s {

		if _, ok := parentthesesMap[v]; ok {
			stack.Push(parentthesesMap[v])
		} else {
			if stack.IsEmpty() {
				return false
			}
			asd, err := stack.Peek()
			if err != nil {
				return false
			}
			// if v != asd {
			// 	return false
			// }
			if v == asd {
				stack.Pop()
			} else if v != asd {
				return false
			}
		}
	}

	if stack.IsEmpty() {
		return true
	} else {
		return false
	}

}
func main() {
	asd := IsValidParentheses("()")
	fmt.Println(asd)
}
