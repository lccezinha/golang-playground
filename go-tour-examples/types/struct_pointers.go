package main

import "fmt"

type Car struct{
  Name string
  Year int
}

func main(){
  car := Car{"Gol", 2005}
  fmt.Println(car)

  p := &car // p recebe um ponteiro para a struct car
  p.Name = "Opala"
  p.Year = 2010
  fmt.Println(car)
}
