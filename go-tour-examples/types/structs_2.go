package main

import "fmt"

type Car struct{
  Name string
  Year int
}

func main(){
  car := Car{"Gol", 2005}
  car.Name = "Monza"

  fmt.Println(car.Name, car.Year)
}
