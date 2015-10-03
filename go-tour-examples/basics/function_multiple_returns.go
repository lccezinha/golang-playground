package main

import "fmt"

// Como essa função retorna dois valores strings eles devem ser passados no final pra
// identificar o retorno de dois tipos, no caso (string, int)
func swap(x int, y string) (string, int) {
  return y, x
}

func main(){
  x, y := swap(5, "Hello")
  fmt.Println(x, y)
}
