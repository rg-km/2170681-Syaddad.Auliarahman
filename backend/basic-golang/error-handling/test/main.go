package main

import "fmt"

// // This function is created to handle
// // the panic occurs in entry function
// // but it does not handle the panic
// // occurred in the entry function
// // because it called in the normal
// // function
// func handlepanic() {

// 	if a := recover(); a != nil {
// 		fmt.Println("RECOVER", a)
// 	}
// }

// // Function
// func entry(lang *string, aname *string) {

// 	// Normal function
// 	handlepanic()

// 	// When the value of lang
// 	// is nil it will panic
// 	if lang == nil {
// 		panic("Error: Language cannot be nil")
// 	}

// 	// When the value of aname
// 	// is nil it will panic
// 	if aname == nil {
// 		panic("Error: Author name cannot be nil")
// 	}

// 	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
// 	fmt.Printf("Return successfully from the entry function")
// }

// // Main function
// func main() {

// 	A_lang := "GO Language"
// 	A_Auth := "Test"
// 	entry(&A_lang, &A_Auth)
// 	fmt.Printf("Return successfully from the main function")
// }

// package main

// import "fmt"

// func main() {
// 	places := make(map[string]int)
// 	places[""] = 1000000000000
// 	places[""] = 100000
// 	fmt.Println(places[""])
// }

// // func addAll(nums ...int) int {
// // 	total := 0
// // 	for _, v := range nums {
// // 		total += v
// // 	}
// // 	return total
// // }

// // func main() {
// // 	points := []int{5, 10, 15}
// // 	fmt.Println(addAll(points...))
// // }

func main() {
	// countryCodes := map[string]string{
	// 	"93":  "afghanistan",
	// 	"374": "Armenia",
	// 	"61":  "Australia",
	// }

	// country := ""
	// ok := false
	//fmt.Println("gina" == "bimo" && 1992 != 1996)

	// var i, j, k int
	// for i = 11; i > 0; i-- {
	// 	for j = 1; j <= i; j++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for k = 11; k > i; k-- {
	// 		fmt.Printf("#")
	// 	}
	// 	fmt.Printf("\n")
	// }

	var i, j, k, l int
	for i = 11; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = 10; k > i; k-- {
			fmt.Printf("*")
		}
		for l = 11; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	// var letters = []string{"a", "b", "c"}
	// var someLetters = make([]string, 2)
	// copy(someLetters, letters)
	// fmt.Println(someLetters)
	// ok := false
	// fmt.Println(ok == false)

	// meaning := map[string]map[string]string{
	// 	"red": map[string]string{
	// 		"element": "fire",
	// 		"feeling": "hot",
	// 	},
	// 	"blue": map[string]string{
	// 		"element": "water",
	// 		"feeling": "cold",
	// 	},
	// }
	// fmt.Println(meaning["blue"]["feeling"])
	// colors := [2]string{"cianjur", "bandung", "jakarta"}
	// fmt.Println(colors[1])

}
