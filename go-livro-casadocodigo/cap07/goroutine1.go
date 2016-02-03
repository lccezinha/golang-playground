package main

import (
  "fmt"
  "time"
)

func takeNap() {
  fmt.Println("takeNap sleeps for 5 seconds...zzz")
  time.Sleep(5 * time.Second)
  fmt.Println("takeNap Close")
}

/*
Goroutines depende da função main pra se manter on, quando a main termina de executar elas também param.
*/
func main() {
  go takeNap()

  fmt.Println("Main sleeps for 3 seconds...zzz")
  time.Sleep(3 * time.Second)
  fmt.Println("Main close")
}