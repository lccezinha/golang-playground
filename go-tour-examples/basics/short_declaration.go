package main

import "fmt"

// Atribuição curta não funciona fora de funções, só funciona dentro de funções
// msg := "Msg"
var msg = "Msg"

func main(){
  var i, j int = 1, 2
  k := 3

  // var c, python, message = false, 1.5, "Message !"
  // OU
  // c, python, message := false, 1.5, "Message !"
  c, python, message := false, 1.5, "Message !"

  fmt.Println(i, j, k)
  fmt.Println(c, python, message)
  fmt.Println(msg)
}
