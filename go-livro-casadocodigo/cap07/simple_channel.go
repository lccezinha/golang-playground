package main

import "fmt"

func sendMessage(c chan string) {
  c <- "Hello Channels" // Envia uma mensagem para o channel
}

func main() {
  c := make(chan string)
  go sendMessage(c)

  message := <- c // Recebe a mensagem do channel

  fmt.Println(message)
}
