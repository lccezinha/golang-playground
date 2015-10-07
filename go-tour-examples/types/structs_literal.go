package main

import "fmt"

type Car struct{
  Model, Factory string
}

func main(){
  a := Car{"Gol", "VW"}
  b := Car{Model: "Monza"}
  c := Car{}
  x := &Car{"Opala", "GM"}

  fmt.Println(a, b, c, x)
}
