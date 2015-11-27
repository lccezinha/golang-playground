package main 

import (
	"fmt"
	"time"
)

// Channels provide a way for two goroutines to communicate with one another and synchronize their execution. 
// Here is an example program using channels:

// Os canais estão sincronizados, quando "pinger" tentar enviar uma mensagem  p/ o canal
// ele irá esperar até que "printer" esteja pronto p/ receber a mensagem.

func pinger(c chan string) {
	for i := 0; ; i ++ {
		c <- "ping" // está colocando a string "ping" no channel c, ou seja, enviando "ping" para o channel
	}
}

func printer(c chan string) {
	for {
		msg := <- c // msg está recebendo a mensagem do channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}