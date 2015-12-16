package main

import "fmt"

type MyInt int

func (i MyInt) IsPositive() bool {
    return i > 0
}

func main() {
    fmt.Println(MyInt(1).IsPositive())
    fmt.Println(MyInt(-1).IsPositive())
}