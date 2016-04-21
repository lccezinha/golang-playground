package main

import (
	"fmt"
	"time"
)

func run(runner string) {
	for i := 1; i <= 10; i++ {
		meters := i * 10
		time.Sleep(2 * time.Second)
		fmt.Printf("%s is running...%dmts \n", runner, meters)
	}
}

func main() {
	bolt := "Bolt"
	powell := "Powell"

	fmt.Println("Start...")
	go run(bolt)
	go run(powell)

	var wait string
	fmt.Scanln(&wait)

	fmt.Println("...End")
}
