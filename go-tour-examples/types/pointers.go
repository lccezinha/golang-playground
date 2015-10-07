package main

import "fmt"

func main(){
  i, j := 42, 210

  p := &i         // p recebe um ponteiro para i
  fmt.Println(*p) // lê o valor de i através do ponteiro p -> 42

  *p = 21         // atribui um novo valor a i através do ponteiro p
  fmt.Println(i, *p) // 21, 21

  x := &j          // x recebe um ponteiro para j
  *x = *x / 105   // ponteiro x (com o valor de j) é divido
  fmt.Println(j)  // 2
}
