//main2.go

package main

import (
   "fmt"
   "github.com/gocarina/gocsv"
   "os"
)

type Record struct {
   Name  string `csv:"name"`
   Gender string `csv:"gender"`
}

func main() {
   // Open the CSV file
   file, err := os.Open("data2.csv")
   if err != nil {
       panic(err)
   }
   defer file.Close()

   // Read the CSV file into a slice of Record structs
   var records []Record
   if err := gocsv.UnmarshalFile(file, &records); err != nil {
       panic(err)
   }

   // Print the records
   for _, record := range records {
       fmt.Printf("Name: %s, Gender: %s\n", record.Name, record.Gender)
   }
}
