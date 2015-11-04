package main 

import "fmt"

func main() {
	const name string = "Luiz"
	fmt.Println(name) // => Luiz
	name = "Cezer" // => erro -> cannot assign to name
}