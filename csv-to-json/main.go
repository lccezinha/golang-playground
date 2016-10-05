package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "encoding/csv"
  "encoding/json"
)

type CommemorativeDate struct {
  Date        string
  Description string
}

type CommemorativeDates []CommemorativeDate
var dates CommemorativeDates

func main () {
  fileToOpen := "./csv-to-json/datas.csv"
  file, err := os.Open(fileToOpen)

  if err != nil {
    panic(err)
  }

  defer file.Close()

  reader := csv.NewReader(file)
  csvInformation, err := reader.ReadAll()

  if err != nil {
    panic(err)
  }

  for _, csvLine := range csvInformation {
    dates = append(dates, CommemorativeDate{Date: csvLine[0], Description: csvLine[1]})
  }

  datesToJson, _ := json.Marshal(dates)
  ioutil.WriteFile("./csv-to-json/output/dates.json", datesToJson, 0644)
  fmt.Println("Done!")
}
