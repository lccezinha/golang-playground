package main

import "fmt"

type ListOfItens []string

func printSlice(slice []string) {
  fmt.Println("Slice", slice)
}

func printList(list ListOfItens) {
  fmt.Println("List", list)
}

func main() {
  list := ListOfItens{"Coca", "Pizza", "Bacon"}
  slice := []string{"Pepsi", "Dogão", "Chesse-Burguer"}

  printSlice([]string(list))
  printList(ListOfItens(slice))
}