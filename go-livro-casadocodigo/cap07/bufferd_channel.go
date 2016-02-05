package main

import "fmt"

func main() {
  c := make(chan int)
  go build(c)

  // fmt.Println(<-c,<-c,<-c,<-c) // Vai dar deadlock, pois o buffer só aceita 3.

  // Dessa forma ocorre a prevenção de causar um deadlock, pois é sempre feita a checagem se o canal está ok

  // for {
  //   value, ok := <-c // Mesma checagem que ocorre para um map.
  //   if ok {
  //     fmt.Printf("Value: %d\n", value)
  //   } else {
  //     break
  //   }
  // }

  // Melhor código V

  for value := range c {
    fmt.Printf("Value: %d\n", value)
  }
}

func build(c chan int) {
  c <- 4
  c <- 5
  c <- 31

  close(c) // Essa função previne de ocorrer o deadlock, pois ela fecha o canal após preencher o buffer
}