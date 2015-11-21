package main 

import "fmt"

func main() {
	value, ok := values()

	fmt.Println(value, ok)
}

// uma função que retorna múltiplos valores deve ter esse retorno (value1, value2) e valores
// separados por vírgula
func values() (string, bool) {
	return "Luiz Cezer", true
}