package main 

import "fmt"

// Múltiplas variáveis
var (
	a = 1
	b = 2
	c = 3

	d, e, f string
)


func main() {
	fmt.Println(a, b, c)

	// Atribuindo valores para as variáveis
	d = "Xunda #1"
	e = "Xunda #2"
	f = "Xunda #3"

	fmt.Println(d, e, f)

	var message string
	message = "Hello World #1"
	fmt.Println(message)

	// Mais utilizado para quando ao criar a variável já se deseja dar algum valor a ela, caso contrário pode-se usar var variable
	message2 := "Hello World #2"
	fmt.Println(message2)

	var message3 string = "Hello World #3"
	fmt.Println(message3)

	// Nomear variável do modo lower camel case -> firstName por exemplo
}