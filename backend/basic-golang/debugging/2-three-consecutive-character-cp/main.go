package main

import (
	"fmt"
)

func main() {
	/*
		Return the start index start from 1 (1 based) of any three consecutive character,
		if not exist return -1

		Example input/output
		www.ruangguru.com -> 1
		helloworld -> -1
	*/
	result := ThreeConsecutiveCharacterPositionCorrect("www.ruangguru.com")
	fmt.Println(result)
}

func ThreeConsecutiveCharacterPosition(word string) int {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] && word[i] == word[i+2] {
			return i
		}
	}
	return -1
}

func ThreeConsecutiveCharacterPositionCorrect(word string) int {
	word += "   "
	empty := " "
	for i := 0; i < len(word); i++ {
		if string(word[i]) == empty {
			return -1
		}
		if word[i] == word[i+1] && word[i] == word[i+2] {
			// fmt.Println(string(word[i]))
			return 1
		}
	}
	return -1 // TODO: replace this
}
