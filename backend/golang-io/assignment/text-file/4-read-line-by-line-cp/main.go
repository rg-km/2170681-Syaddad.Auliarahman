package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {
	// return nil // TODO: replace this
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}

	return nil
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	// return nil // TODO: replace this
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			keyValue := strings.Split(line, ",")
			dataMap[keyValue[0]] = keyValue[1]
		}
	}

	return nil
}
