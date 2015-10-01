package main

import "fmt"

// add(x int, y int) -> como os dois são do mesmo tipo, pode ser omitido e trocado por
// add(x, y int)
func add(x, y int) int {
  return x + y
}

func main(){
  fmt.Println("Soma é: ", add(1, 6))
}
