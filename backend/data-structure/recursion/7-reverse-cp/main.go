// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	// TODO: answer here
	if len(st) < 0 {
		return str
	}
	// fmt.Println(st[:len(st)-3])
	for len(st) > 0 {
		str += st[len(st)-1]
		st = st[:len(st)-1]
	}

	return str
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
