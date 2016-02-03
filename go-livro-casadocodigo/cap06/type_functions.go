package main

import "fmt"

type Aggregable func(n, m int) int

func Aggregator(values []int, start int, fn Aggregable) int {
  accum := start

  for _, value := range values {
    accum = fn(value, accum)
  }

  return accum
}

func CalculateSum(values []int) int {
  sum := func(n, m int) int {
    return n + m
  }

  return Aggregator(values, 0, sum)
}

func CalculateMult(values []int) int {
  mult := func(n, m int) int {
    return n * m
  }

  return Aggregator(values, 1, mult)
}

func main() {
  values := []int{3, 5, -4, 8, 10, 8, 5}

  fmt.Println(CalculateSum(values))
  fmt.Println(CalculateMult(values))
}