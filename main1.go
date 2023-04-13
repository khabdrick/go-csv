//main1.go
package main

import (
    "os"

    "github.com/gocarina/gocsv"
)

type Person struct {
    Name   string `csv:"name"`
    Age    int    `csv:"age"`
    Gender string `csv:"gender"`
}

func main() {
    file, err := os.Create("data2.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    people := []*Person{
        {"Alice", 25, "Female"},
        {"Bob", 30, "Male"},
        {"Charlie", 35, "Male"},
    }

    if err := gocsv.MarshalFile(&people, file); err != nil {
        panic(err)
    }
}