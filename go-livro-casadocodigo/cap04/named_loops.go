package main

import "fmt"

func main() {
  var i int

  loop:
  for i = 0; i < 10; i++ {
    fmt.Printf("for i = %d\n", i)

    switch i {
    case 2, 3:
      fmt.Printf("Broke switch, i == %d\n.", i)
      break
    case 5:
      fmt.Printf("Broke loop, i == 5.")
      break loop
    }
  }

  fmt.Println("End.")
}