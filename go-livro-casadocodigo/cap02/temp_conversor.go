package main

import (
  "fmt"
  "strconv"
  "os"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("Uso: conversor <valores> <unidade>")
    os.Exit(1)
  }

  lastPosition := len(os.Args) - 1
  unitOrig := os.Args[lastPosition]
  valuesOrig := os.Args[1 : lastPosition]

  var unitDestiny string

  if unitOrig == "celsius" {
    unitDestiny = "fahranheit"
  } else if unitOrig == "km" {
    unitDestiny = "milhas"
  } else {
    fmt.Printf("%s invalid!", unitOrig)
    os.Exit(1)
  }

  for index, value := range valuesOrig {
    valueOrig, err := strconv.ParseFloat(value, 64)

    if err != nil {
      fmt.Printf("Value %s in position %d is not valid!", value, index)
      os.Exit(1)
    }
    var valueDestiny float64

    if unitOrig == "celsius" {
      valueDestiny = valueOrig * 1.8 + 32
    } else {
      valueDestiny = valueOrig * 1.60934
    }

    fmt.Printf("%.2f %s = %.2f %s \n", valueOrig, unitOrig, valueDestiny, unitDestiny)
  }
}