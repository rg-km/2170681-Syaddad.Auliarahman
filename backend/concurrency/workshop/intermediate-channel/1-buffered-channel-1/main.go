package main

import "fmt"

//buferred channel tidak blocking jika buffer belum penuh
// func test(count int, input chan string) (string, int) {
// 	var x string
// 	fmt.Scanf("input :", &x)
// 	// input <- x
// 	count++
// 	return x, count
// }

func main() {
	input := make(chan string, 4)
	// capx := cap(input)
	// count := 0
	// for i := 0; i < capx; i++ {
	// 	if count <= capx {
	// 		y, count1 := test(count, input)
	// 		input <- y
	// 		count = count1
	// 	}

	// }
	var arr []string
	var x string
	for i := 0; i < cap(input); i++ {
		fmt.Scanf("Input: ", &x)
		arr = append(arr, x)
		input <- arr[i]
	}
	close(input)

	// input <- "halo"
	// input <- "dari buffered"

	for i := 0; i < cap(input); i++ {
		//fmt.Println(cap(input))
		fmt.Println(<-input)
	}
}
