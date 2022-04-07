package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CSVRead() {
	f, err := os.Open("numbers.csv")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(f)

	r := csv.NewReader(f)

	//for {
	record, err := r.Read()
	if err == io.EOF {
		fmt.Println("\nthis is the End Of File")
		//break
	}

	if err != nil {
		log.Fatal(err)
	}
	for value := range record {
		fmt.Print(value)
		fmt.Printf("%s ", record[value])
	}
	//}
	f.Close()
	//defer f.Close()
}
