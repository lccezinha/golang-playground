package main

import "fmt"

func minus(n1, n2 int) int {
    return n1 - n2
}

func main() {
    n1 := 20
    n2 := 10

    fmt.Println(minus(n1, n2))
}