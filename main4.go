//main4.go
package main

import (
    "bytes"
    "encoding/json"
    "github.com/yukithm/json2csv"
    "log"
    "os"
)

func main() {
    b := &bytes.Buffer{}
    wr := json2csv.NewCSVWriter(b)
    j, _ := os.ReadFile("input.json")
    var x []map[string]interface{}

    // unMarshall json
    err := json.Unmarshal(j, &x)
    if err != nil {
        log.Fatal(err)
    }

    // convert json to CSV
    csv, err := json2csv.JSON2CSV(x)
    if err != nil {
        log.Fatal(err)
    }

    // CSV bytes convert & writing...
    err = wr.WriteCSV(csv)
    if err != nil {
        log.Fatal(err)
    }
    wr.Flush()
    got := b.String()

    //Following line prints CSV
    // println(got)

    // create file and append if you want
    createFileAppendText("output.csv", got)
}

//
func createFileAppendText(filename string, text string) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }
}