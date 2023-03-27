package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

func main() {
	// Open the CSV file
	file, err := os.Open("semi.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create the CSV reader with semicolon as delimiter
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read the headers
	headers, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// If no header was provided, use the first line of the CSV file as the header
	if len(headers) == 0 {
		if headers, err = reader.Read(); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Read the data into a slice of maps
	var data []map[string]interface{}
	decoder, err := csvutil.NewDecoder(csv.NewReader(file), headers...)

	for {
		var record map[string]interface{}
		if err := decoder.Decode(&record); err != nil {
			break
		}
		data = append(data, record)
	}

	// Print the data
	for _, record := range data {
		fmt.Println(record)
	}
}
