package main

import "errors"
import "fmt"

func main() {
  err := errors.New("Xunda")

  fmt.Println(err)
}