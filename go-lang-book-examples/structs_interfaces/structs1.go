package main

import "fmt"

type Circle struct {
	x, y, z float64
}

func main() {
	// var c Circle
	// c.x = 10.0
	// c.y = 11.0
	// c.z = 12.0

	// c := new(Circle) // Cria como um ponteiro
	// c.x = 10.0
	// c.y = 11.0
	// c.z = 12.0

	// c := Circle{x: 1.0, y: 2.0, z: 3.0}

	c := Circle{1.0, 2.0, 3.0}
	fmt.Println(c)
	fmt.Println(c.x, c.y, c.z)
}