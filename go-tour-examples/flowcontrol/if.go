package main

import "fmt"

func double(x int) string {
  if (x * 2) >= 20{
    return "X MAIOR QUE 20"
  }

  return "deu ruim"
}

func main(){
  number := 5
  number_2 := 10

  fmt.Println(double(number))
  fmt.Println(double(number_2))
}
