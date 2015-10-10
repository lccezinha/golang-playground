package main

import "fmt"

type Car struct{
  Model, Factory string
}

func main(){
  a := Car{"Gol", "VW"}
  b := Car{Model: "Monza"} // Factory fica em branco
  c := Car{} // Tudo em branco
  x := &Car{"Opala", "GM"}

  fmt.Println(a, b, c, x)
}
