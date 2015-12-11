package main 

import "fmt"

func main() {
    // names := []string{}
    // names[0] = "Cezinha" -> erro
    // fmt.Println(names)

    // As explained above, a slice is seating on top of an array, in this case, 
    // the array is empty and the slice can’t set a value in the referred array. 
    // There is a way to do that though, and that is by using the append function:

    names := []string{}
    names = append(names, "Luiz", "Cezer")
    fmt.Println(names)
}