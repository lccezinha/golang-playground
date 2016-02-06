package main

import (
  "fmt"
  "math"
  "sync"
  "time"
)

func calculate(base float64, control *sync.WaitGroup) {
  defer control.Done()
  n := 0.0

  for i := 0; i < 10000000; i++ {
    n += base / math.Pi* math.Sin(2)
  }

  fmt.Println(n)
}

func main() {
  start := time.Now()
  var control sync.WaitGroup
  control.Add(3)

  go calculate(9.37, &control)
  go calculate(6.55, &control)
  go calculate(12.21, &control)

  control.Wait()
  fmt.Printf("Finished in %s. \n", time.Since(start))
}