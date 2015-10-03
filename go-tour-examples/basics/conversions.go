package main

import "fmt"

// Atribuição para tipos diferentes são EXPLÍCITAS

func main(){
  i := 10
  var z float32 = 15.2

  fmt.Println(i, z)

  var x int = int(z)
  var y float32 = float32(i)

  fmt.Println(x, y)
}

