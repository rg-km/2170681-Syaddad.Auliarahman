package main

import (
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {
	// return nil // TODO: replace this
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	len, err := file.WriteString(fileData)
	if err != nil {
		return err
	}

	log.Printf("File Name: %s\n", file.Name())
	log.Printf("Length: %d bytes\n", len)
	return nil
}
