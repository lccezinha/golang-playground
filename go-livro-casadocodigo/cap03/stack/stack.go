package main

import (
  "errors"
  "fmt"
)

type Stack struct {
  elements []interface{} // "Interface vazia" -> descreve uma interface que implemente ao menos 0 métodos, qualquer tipo em Go obedece essa condição, por isso pode aceitar qualquer item.
}

func (stack Stack) Empty() bool {
  return len(stack.elements) == 0
}

func (stack Stack) Size() int {
  return len(stack.elements)
}

func (stack *Stack) Push(value interface{}) {
  stack.elements = append(stack.elements, value)
}

func (stack *Stack) Pop() (interface{}, error) {
  if stack.Empty() {
    return nil, errors.New("Empty Stack")
  }

  value := stack.elements[stack.Size() - 1]
  stack.elements = stack.elements[:stack.Size() - 1]

  return value, nil
}

func main() {
  stack := Stack{}

  fmt.Println("Empty ?", stack.Empty())
  fmt.Println("Size", stack.Size())

  stack.Push(1)
  stack.Push("Xunda")
  stack.Push(3.2)

  fmt.Println("Empty ?", stack.Empty())
  fmt.Println("Size", stack.Size())

  for !stack.Empty() {
    value, _ := stack.Pop()

    fmt.Println("Pop", value)
    fmt.Println("Empty ?", stack.Empty())
    fmt.Println("Size", stack.Size())
  }

  _, err := stack.Pop()
  if err != nil {
    fmt.Println("Error:", err)
  }
}