package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

type Questions struct {
	quest string
	ans   string
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here
	f, err := os.Open("questions.csv")
	if err != nil {
		return data, err
	}
	defer f.Close()
	rdr := csv.NewReader(f)

	if _, err := rdr.Read(); err != nil {
		return data, err
	}

	data1, err := rdr.ReadAll()
	if err != nil {
		return data, err
	}
	for _, record := range data1 {
		data[record[0]] = record[1]
	}
	return data, nil

}
