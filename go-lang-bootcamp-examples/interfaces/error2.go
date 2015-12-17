package main

import "fmt"

type NumberError int

func (e NumberError) Error() string {
    return fmt.Sprintf("Number must be positive, but number is negative. Number: %d", e)
}

func Double(n int) (int, error) {
    if n <= 0 {
        return 0, NumberError(n)
    }

    result := n * 2
    return result, nil
}

func main() {
    fmt.Println(Double(2))
    fmt.Println(Double(-12))
}