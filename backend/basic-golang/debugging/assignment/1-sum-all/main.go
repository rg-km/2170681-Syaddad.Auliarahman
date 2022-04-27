package main

import (
	"fmt"
)

func main() {
	arr := []int{12, 5, 10, 1, 3}

	res := SumAllCorrect(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := SumAllCorrect(arr)
	// fmt.Println(resCorrect)
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	res := 0
	for _, val := range arr {
		//fmt.Println(val)
		res += val
	}
	return res

	//return 0 // TODO: replace this
}
