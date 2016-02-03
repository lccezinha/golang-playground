package main

import (
  "fmt"
  "time"
)

func GenerateFibonacci(n int) func() {
  return func() {
    a, b := 0, 1

    fib := func() int {
      a, b = b, b + a

      return a
    }

    for i := 0; i < n; i++ {
      fmt.Printf("%d" , fib())
    }
  }
}

func Timer(function func()) {
  start := time.Now()

  function()

  fmt.Printf("\nTime to execute: %s\n\n", time.Since(start))
}

func main() {
  Timer(GenerateFibonacci(8))
  Timer(GenerateFibonacci(48))
  Timer(GenerateFibonacci(88))
}