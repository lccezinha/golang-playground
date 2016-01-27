package main

import (
  "fmt"
  "os"
  "strconv"
)

func checkError(err error, numberString string) {
  if err != nil {
    fmt.Printf("%s not valid! \n", numberString)
    os.Exit(1)
  }
}

func quicksort(numbers []int) []int {
  if len(numbers) <= 1 {
    return numbers
  }

  n := make([]int, len(numbers))
  copy(n, numbers)

  indexCentral := len(n) / 2
  itemCentral := n[indexCentral]
  n = append(n[:indexCentral], n[indexCentral+1:]...)

  majors, minors := separate(n, itemCentral)

  return append(
    append(quicksort(minors), itemCentral),
    quicksort(majors)...)
}

func separate(numbers []int, itemCentral int) (minors []int, majors[] int) {
  for _, n := range numbers {
    if n <= itemCentral {
      minors = append(minors, n)
    } else {
      majors = append(majors, n)
    }
  }

  return minors, majors
}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("No args!")
    os.Exit(1)
  }

  entries := os.Args[1:]
  numbers := make([]int, len(entries))

  for index, numberString := range entries {
    number, err := strconv.Atoi(numberString) // Converte String -> Inteiro
    checkError(err, numberString)
    numbers[index] = number
  }

  fmt.Println(quicksort(numbers))
}