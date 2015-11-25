package main 

import "fmt"

// func zero(number int) {
// 	number = 0
// }

// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // Nesse caso x continua sendo 5, por que o number = 0, utiliza apenas uma copia de X
// 	// e não a variável de vdd
// }

func zero(pointerToNumber *int) {
	*pointerToNumber = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // Nesse caso muda p/ 0, por que foi utilizada a variável e não uma cópia
}

// Pointeiro é representado por um * seguido pelo tipo da variável

// *pointerToNumber int = 1 --> guarde o inteiro 1 no local de mémoria onde está a referência p/ pointerToNumber
