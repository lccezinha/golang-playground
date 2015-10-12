package main

import "fmt"

func main(){
  m := make(map[string]int)

  m["Pergunta"] = 42
  fmt.Println("The value", m["Pergunta"])

  m["Pergunta"] = 48
  fmt.Println("The value", m["Pergunta"])

  delete(m, "Pergunta")
  fmt.Println("The value", m["Pergunta"])

  m["Pergunta"] = 48
  value, ok := m["Pergunta"]
  fmt.Println("The value", value, " Present?", ok)
}
