package main 

import "fmt"

func main() {
    bands := []string{"Metallica", "Trivium", "Slayer"}
    fmt.Println(len(bands)) // => 3

    otherBands := make([]string, 10)
    fmt.Println(len(otherBands))
}