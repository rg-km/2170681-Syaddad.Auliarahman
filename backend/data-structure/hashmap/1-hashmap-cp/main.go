package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{
		{name: "Bob", age: 21},
		{name: "Sam", age: 28},
		{name: "Ann", age: 21},
		{name: "Joe", age: 22},
		{name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	res := make(map[int]int)
	for _, age := range people {
		intAge := int(age.age)
		if _, ok := res[intAge]; ok {
			res[intAge] += 1
		} else {
			res[intAge] = 1
		}

	}
	return res // TODO: replace this
}

func FilterByAge(people []Person, age int) []Person {
	var result []Person
	for _, ageMap := range people {
		intAge := int(ageMap.age)
		if intAge == age {
			result = append(result, ageMap)
		}
	}
	return result // TODO: replace this
}
