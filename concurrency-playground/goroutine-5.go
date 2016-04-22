package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func run(runner string, winner chan string, waitGroup *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		meters := i * 10

		delay := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(delay)

		fmt.Printf("%s is running...%dmts \n", runner, meters)

		if meters == 100 {
			winner <- runner
		}
	}

	defer waitGroup.Done()
}

func main() {
	winner := make(chan string, 5)
	var waitGroup sync.WaitGroup

	fmt.Println("Start...")

	waitGroup.Add(5)
	go run("Bolt", winner, &waitGroup)
	go run("Powell", winner, &waitGroup)
	go run("The Flash", winner, &waitGroup)
	go run("Gzuis", winner, &waitGroup)
	go run("Oba Oba Martins", winner, &waitGroup)
	waitGroup.Wait()

	fmt.Println("\n---\n")
	fmt.Println("1 -", <-winner)
	fmt.Println("2 -", <-winner)
	fmt.Println("3 -", <-winner)
	fmt.Println("4 -", <-winner)
	fmt.Println("5 -", <-winner)
	fmt.Println("\n...End")
}
