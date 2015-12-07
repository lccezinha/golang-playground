package main

import "fmt"

func sum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	n1 := 10
	n2 := 20

	fmt.Println(sum(n1, n2))
}