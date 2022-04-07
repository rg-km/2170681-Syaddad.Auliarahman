package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// TODO: answer here
	siWriter := csv.NewWriter(csvFile)
	defer siWriter.Flush()

	for _, record := range records {
		row := []string{record.Name, strconv.Itoa(record.Price)}
		if err := siWriter.Write(row); err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
