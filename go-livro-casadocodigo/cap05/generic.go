package main

import "fmt"

func (list *GenericList) RemoveItem(index int) interface{} {
  l := *list
  removed := l[index]
  *list = append(l[0:index], l[index+1:]...)

  return removed
}

func (list *GenericList) RemoveStart() interface{} {
  return list.RemoveItem(0)
}

func (list *GenericList) RemoveEnd() interface{} {
  return list.RemoveItem(len(*list) - 1)
}

type GenericList []interface{}

func main() {
  list := GenericList{1, "Café", 42, true, 30, "Bola", 1.4, false}

  fmt.Println("List :", list)
  fmt.Println("RemoveItem :", list.RemoveItem(2), list)
  fmt.Println("RemoveEnd :", list.RemoveEnd(), list)
  fmt.Println("RemoveStart :", list.RemoveStart(), list)
}