package main

import "fmt"

/*
  Valores retornados por uma função, quando nomeados no momento da declaração da função
  ficam disponíveis para serem utilizados dentro do escopo da função,
  dessa forma é possível só utilizar `return` sem argumentos.
*/

func finalValues(price float64) (finalPrice, doublePrice float64) {
  tax := 1.5
  doubleTax := tax * 2

  finalPrice = price * tax
  doublePrice = price * doubleTax

  return
}

func otherValues(price float64) (float64, float64) {
  tax := 1.3
  doubleTax := tax * 2

  finalPrice := price * tax
  doublePrice := price * doubleTax

  // return ./name_return_values.go:28: not enough arguments to return
  return finalPrice, doublePrice
}

func main() {
  finalPrice, doublePrice := finalValues(11.5)
  fp, dp := otherValues(11.5)

  fmt.Printf("%.2f - %.2f \n", finalPrice, doublePrice)
  fmt.Printf("%.2f - %.2f \n", fp, dp)
}