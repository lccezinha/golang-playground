package main

import "fmt"

func main() {
    people := map[string]int{
        "Luiz Cezer": 25,
        "Cezer Luiz": 52,
        "Cezinha Cezer": 11,
    }

    fmt.Printf("%#v", people)
}