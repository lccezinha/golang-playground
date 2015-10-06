package main

import "fmt"

func main(){
  sum := 1

  for sum <= 10 {
    fmt.Println("Atual:", sum)
    sum += sum
  }

  fmt.Println("Total:", sum)
}
