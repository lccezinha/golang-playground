package main

import "fmt"

type ListOfItens []string // Tipos customizados podem ser estendidos, ao contrário dos tipos padrão.

func (list ListOfItens) Categorize() ([]string, []string, []string){
  var veg, meat, other []string

  for _, e := range list {
    switch e {
    case "Alface", "Tomate":
      veg = append(veg, e)

    case "Carne", "Bacon":
      meat = append(meat, e)

    default:
      other = append(other, e)
    }
  }

  return veg, meat, other
}

func main() {
  list := make(ListOfItens, 4)

  list[0] = "Coca"
  list[1] = "Pizza"
  list[2] = "Bacon"
  list[3] = "Tomate"

  fmt.Println(list)

  veg, meat, other := list.Categorize()

  fmt.Println("Veg", veg)
  fmt.Println("Meat", meat)
  fmt.Println("Other", other)
}