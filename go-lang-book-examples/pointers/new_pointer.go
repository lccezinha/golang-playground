package main 

import "fmt"

func value(pointerToNumber *int) {
	*pointerToNumber = 10
}

// new() aloca o espaço em memória necessário p/ alocar o tipo de variável
func main() {
	number := new(int)
	value(number)
	fmt.Println(*number)
}