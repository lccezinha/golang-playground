package main

import (
  "fmt"
  "os"
)

func checkError(err error, string numberString) {
  if err != nil {
    fmt.Printf("%s not valid! \n", numberString)
    os.Exit(1)
  }
}

func quicksort()

func main() {
  entry := os.Args[1:]
  numbers := make([]int, len(entry))

  for index, numberString := entry {
    number, err := strconv.Atoi(numberString) // Converte String -> Inteiro

    checkError(err, numberString)

    numbers[index] = number
  }

  fmt.Println(quicksort(numbers))
}