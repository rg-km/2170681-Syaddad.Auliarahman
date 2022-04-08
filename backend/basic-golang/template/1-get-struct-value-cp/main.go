package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	// TODO: answer here
	Nama         string
	ScoreAverage float64
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
	std := Student{
		Nama:         "Rogu",
		ScoreAverage: 7.8,
	}
	tmp1 := template.New("Template_1")

	tmp1, err := tmp1.Parse("Hello {{.Nama}},\nNilai rata rata kamu {{.ScoreAverage}}")
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
