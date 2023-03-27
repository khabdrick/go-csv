package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("Data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Print the CSV data
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}

	// Write the CSV data
	file2, err := os.Create("Data1.csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()
// this defines the header value and data values for the new csv file
	headers := []string{"name", "age", "gender"}
	Data1 := [][]string{
		{"Alice", "25", "Female"},
		{"Bob", "30", "Male"},
		{"Charlie", "35", "Male"},
	}

	writer.Write(headers)
	for _, row := range Data1 {
		writer.Write(row)
	}
}