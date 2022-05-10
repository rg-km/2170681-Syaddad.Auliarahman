package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	// for _, val1 := range arr1{
	var res []string
	found := false
	// }
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			fmt.Println(string(arr1[i]), string(arr2[j]))
			if arr1[i] == arr2[j] {
				found = true
				res = append(res, string(arr1[i]))
			}
		}
	}
	if found == false {
		return []string{}, fmt.Errorf("no match")
	}
	return res, nil

	// TODO: replace this
}
