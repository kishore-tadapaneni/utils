package main

import (
	"fmt"
	"log"

	"github.com/kishore-tadapaneni/utils/csvutil"
	"github.com/kishore-tadapaneni/utils/jsonutil"
)

func main() {
	inputFilePath := "input.csv"
	outputFilePath := "ouput.json"

	employees, err := csvutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}
	if err := jsonutil.WriteFile(outputFilePath, employees); err != nil {
		log.Fatalf("Error writing JSON data: %v", err)
	}
	fmt.Println("Conversion completed. Results written to output.json.")
}
