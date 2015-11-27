package main 

import (
	"fmt"
	"time"
)

//

// Canais com essa assinatura são bi-direcionais, podem maandar ou receber mensagens
/* 
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // está colocando a string "ping" no channel c, ou seja, enviando "ping" para o channel
	}
}
*/

// Channels provide a way for two goroutines to communicate with one another and synchronize their execution. 
// Here is an example program using channels:

// Os canais estão sincronizados, quando "pinger" tentar enviar uma mensagem  p/ o canal
// ele irá esperar até que "printer" esteja pronto p/ receber a mensagem.

// Envia mensagem para o channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping" // está colocando a string "ping" no channel c, ou seja, enviando "ping" para o channel
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// Recebe mensagem do channel
func printer(c <-chan string) {
	for {
		msg := <- c // msg está recebendo a mensagem do channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}