package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func startAlarm() {
	fmt.Println("\nArming Alarm...")
	time.Sleep(time.Duration(3) * time.Second) //60 seconds
}

func getReady(who string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	timeSpent := rand.Intn(30) + 60

	fmt.Printf("%s started getting ready\n", who)
	time.Sleep(time.Duration(timeSpent) * time.Second)
	fmt.Printf("%s spent %d seconds to get ready\n", who, timeSpent)
}

func putShoes(who string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	timeSpent := rand.Intn(10) + 35

	fmt.Printf("%s started putting on shoes\n", who)
	time.Sleep(time.Duration(timeSpent) * time.Second)
	fmt.Printf("%s finished to put shoes is %d seconds\n", who, timeSpent)
}

func prepareAlarm(armFinished chan<- struct{}) {
	fmt.Println("\nArming alarm..")
	time.Sleep(60 * time.Second)
	close(armFinished)
}

func main() {
	bob, alice := "Bob", "Alice"
	var waitGroup sync.WaitGroup

	fmt.Println("Let's walk...\n")

	waitGroup.Add(2)
	go getReady(bob, &waitGroup)
	go getReady(alice, &waitGroup)
	waitGroup.Wait()

	// armStarted := make(chan struct{})
	armFinished := make(chan struct{})
	go prepareAlarm(armFinished)
	// <-armStarted // Recebe o sinal de close(armStarted)

	waitGroup.Add(3)
	go func() {
		fmt.Println("Alarm is counting down...")
		waitGroup.Done()
	}()

	go putShoes(bob, &waitGroup)
	go putShoes(alice, &waitGroup)
	waitGroup.Wait()

	fmt.Println("Exiting and locking door.")
	<-armFinished
	fmt.Println("Alarm is ARMED!")
}

// Every morning, Alice and Bob go for a walk, and being creatures of habit,
// they follow the same routine every day.

// First, they both prepare, grabbing sunglasses, perhaps a belt, closing open windows,
// turning off ceiling fans, and pocketing their phones and keys.

// Once they’re both ready, which typically takes each of them between 60 and 90 seconds,
// they arm the alarm, which has a 60 second delay.

// While the alarm is counting down, they both put on their shoes,
// a process which tends to take each of them between 35 and 45 seconds.

// Then they leave the house together and lock the door, before the alarm has finished its countdown.

// Write a program to simulate Alice and Bob’s morning routine.

// Here’s some sample output from running a solution to this problem.

// Let's go for a walk!
// Bob started getting ready
// Alice started getting ready
// Alice spent 72 seconds getting ready
// Bob spent 76 seconds getting ready

// Arming alarm.

// Bob started putting on shoes
// Alarm is counting down.
// Alice started putting on shoes
// Alice spent 37 seconds putting on shoes
// Bob spent 39 seconds putting on shoes

// Exiting and locking the door.
// Alarm is armed.
