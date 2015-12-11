package main 

import "fmt"

func main() {
    names := []string{}
    names = append(names, "Luiz", "Cezer")
    fmt.Println(names)

    otherNames := []string{"Cezinha", "Xunda"}
    names = append(names, otherNames...)
    fmt.Println(names)
}