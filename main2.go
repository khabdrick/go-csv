//main2.go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    // Open the CSV file for appending
    file, err := os.OpenFile("Data1.csv", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Create a CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write a new row to the CSV file
    row := []string{"David", "30", "Male"}
    err = writer.Write(row)
    if err != nil {
        panic(err)
    }
}