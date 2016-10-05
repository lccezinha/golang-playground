package main

import (
  "fmt"
  "os"
  "encoding/csv"
  "encoding/json"
)

type CommemorativeDate struct {
  Date        string `json:"date"`
  Description string `json:"description"`
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
  jsonFile, err := os.Create("./csv-to-json/output/dates.json")

  if err != nil {
    panic(err)
  }
  defer jsonFile.Close()

  jsonFile.Write(datesToJson)

  fmt.Println("Done!")
}
