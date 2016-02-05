package main

import "fmt"

func main() {
  c := make(chan int)
  go build(c)

  value := <-c

  fmt.Println(value)
}

// Dessa forma só é possível enviar dados para o channel
func build(c chan<- int) {
  c <- 4
  c <- 5
  c <- 31

  close(c) // Essa função previne de ocorrer o deadlock, pois ela fecha o canal após preencher o buffer
}