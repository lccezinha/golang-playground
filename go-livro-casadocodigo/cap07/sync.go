package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

func execute(control *sync.WaitGroup) {
  defer control.Done()

  duration := time.Duration(1 + rand.Intn(5)) * time.Second
  fmt.Printf("Sleeping by %s ... \n", duration)
  time.Sleep(duration)
}

func main() {
  start := time.Now()
  rand.Seed(start.UnixNano())

  var control sync.WaitGroup

  for i := 0; i < 5; i++ {
    control.Add(1)
    go execute(&control)
  }

  control.Wait()

  fmt.Printf("Finished em %s. \n", time.Since(start))
}