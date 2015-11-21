package main

import "fmt"

func main() {
	fmt.Println(values(1))
	fmt.Println(values(2))
}

func values(number int) (int, bool) {
	even := false

	if number % 2 == 0 {
		even = true
	}

	half := number / 2

	return half, even
}