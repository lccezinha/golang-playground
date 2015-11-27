package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

// A goroutine is a function that is capable of running concurrently with other functions. 
// To create a goroutine we use the keyword go followed by a function invocation
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}