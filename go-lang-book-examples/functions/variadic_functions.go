package main

import "fmt"

func main() {
	fmt.Println(values(1, 4, 5))
}

// Indica que o argumento é uma lista de 0 a N valores, igual *args do ruby.
func values(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
