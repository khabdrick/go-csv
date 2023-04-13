//main5.go
package main

import (
    "encoding/csv"
    "encoding/json"
    "log"
    "os"
    "strconv"
)

func main() {
    // Open the CSV file
    file, err := os.Open("data1.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create the CSV reader
    csvReader := csv.NewReader(file)

    // Read the CSV headers
    headers, err := csvReader.Read()
    if err != nil {
        log.Fatal(err)
    }

    // Read the CSV data rows
    var data []map[string]interface{}
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }
        // Convert the row values to the appropriate types
        m := make(map[string]interface{})
        for i, val := range row {
            f, err := strconv.ParseFloat(val, 64)
            if err == nil {
                m[headers[i]] = f
                continue
            }
            b, err := strconv.ParseBool(val)
            if err == nil {
                m[headers[i]] = b
                continue
            }
            m[headers[i]] = val
        }
        data = append(data, m)
    }

    // Encode the JSON data and write it to stdout
    encoder := json.NewEncoder(os.Stdout)
    if err := encoder.Encode(data); err != nil {
        log.Fatal(err)
    }
}
