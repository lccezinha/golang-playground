package main

import "fmt"

func main() {
  p := []int{2, 3, 5, 8, 11, 13}
  fmt.Println("p ==", p)

  //pega do 1 até o 4 item do ~array~ (1:4) são os indexes
  // [X:Y]
  // X = A partir do item
  // Y = Até o index informado - 1
  fmt.Println("p[1:4] ==", p[1:4])

  //pega até o terceiro
  fmt.Println("p[:3] ==", p[:3])

  //pega do quarto em diante
  fmt.Println("p[4:] ==", p[4:])

  fmt.Println("p[0:1]", p[0:2])
}
