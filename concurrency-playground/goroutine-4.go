package main

import (
	"fmt"
	"math/rand"
	"time"
)

func run(runner string, winner chan string) {
	for i := 1; i <= 10; i++ {
		meters := i * 10

		delay := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(delay)

		fmt.Printf("%s is running...%dmts \n", runner, meters)

		if meters == 100 {
			winner <- runner
		}
	}
}

func main() {
	winner := make(chan string, 5)

	fmt.Println("Start...")

	go run("Bolt", winner)
	go run("Powell", winner)
	go run("The Flash", winner)
	go run("Gzuis", winner)
	go run("Oba Oba Martins", winner)

	time.Sleep(1 * time.Second)

	fmt.Println("\n---\n")
	fmt.Println("1 -", <-winner)
	fmt.Println("2 -", <-winner)
	fmt.Println("3 -", <-winner)
	fmt.Println("4 -", <-winner)
	fmt.Println("5 -", <-winner)
	fmt.Println("...End")
}
