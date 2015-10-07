package main

import "fmt"

type Car struct{
  name string
  year int
}

func main(){
  fmt.Println(Car{"Gol", 2005})
}
