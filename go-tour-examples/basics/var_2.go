package main

import "fmt"

// var some, any int = 1, 3
// duck type

// Se um inicializador está presente, o tipo pode ser omitido; a variável terá o tipo do inicializador.
var i, j int = 1, 2

func main(){
  var ruby, python, message = false, true, "Message!"

  fmt.Println(i, j, ruby, python, message)
}
