package main

import "fmt"

func main(){
  sum := 1

  for i := 1; i <= 10; i++ {
    fmt.Println("Atual:", i)
    sum += i
  }

  fmt.Println("Total:", sum)
}
