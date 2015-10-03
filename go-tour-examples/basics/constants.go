package main

import "fmt"

// constantes utilizando const ao invés de var
// não podem utilizar a sintaxe curta (:=)

const Pi = 3.14

func main(){
  const World = "Hello"

  fmt.Println("Pi", Pi)
  fmt.Println("Hello", World)

  const Truth = true

  fmt.Println("Go Rules ?", Truth)
}
