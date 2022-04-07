package main

import (
	"fmt"
	"log"
)

func main() {
	CSVRead()
	data, err := CSVReadAll()
	if err != nil {
		fmt.Println(" ", err)
		log.Fatal(err)

	}
	//fmt.Println(data)

	for x, record := range data {
		fmt.Println(x)
		user := Person{
			FirstName: record[0],
			LastName:  record[1],
			Domicile:  record[2],
		}
		fmt.Println(user)
	}
}
