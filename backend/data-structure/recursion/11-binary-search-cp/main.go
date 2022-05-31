package main

import (
	"fmt"
	"time"
)

func main() {
	numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := int64(9)
	fmt.Println(BinarySearch(numList, key))
}

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1
	time.Sleep(time.Second * 1)
	fmt.Println(low, high)
	if low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if numList[mid] == key {
			fmt.Println(numList[mid])
			return 1
		}
		if numList[mid] > key {
			fmt.Println(numList)
			numList = numList[low:mid]
			BinarySearch(numList[low:mid], key)
		}
		fmt.Println(numList)
		numList = numList[1:]
		return BinarySearch(numList[1:], key)

		// TODO: answer here
	}
	return 0
}
