package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

// Ao usar 'defer' ele garante que o método será executado por último.

// Geralmente  utilizado quando o recurso após executar, precisa liberar mémoria
func main() {
	first()
	defer second()
	first()
}