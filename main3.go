//main3.go
package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    jsonData, err := ioutil.ReadFile("input.json")
    if err != nil {
        log.Fatal(err)
    }

    var data []map[string]interface{}
    err = json.Unmarshal(jsonData, &data)
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create("output.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)

    header := make([]string, 0, len(data[0]))
    for key := range data[0] {
        header = append(header, key)
    }

    err = writer.Write(header)
    if err != nil {
        log.Fatal(err)
    }

    for _, row := range data {
        values := make([]string, 0, len(row))
        for _, value := range row {
            values = append(values, fmt.Sprintf("%v", value))
        }
        err = writer.Write(values)
        if err != nil {
            log.Fatal(err)
        }
    }

    writer.Flush()
}