package main

import (
    "fmt"
    "sort"
)

func main() {
    letters := []string{ "f", "g", "a", "c", "e", "z", "b" }
    result := sort.StringsAreSorted(letters)
    fmt.Println(result)

    sort.Strings(letters)
    fmt.Println(letters)

    result = sort.StringsAreSorted(letters)
    fmt.Println(result)
    
    fmt.Println("---")

    numbers := []int{ 10, 40, 20, 5 }
    result = sort.IntsAreSorted(numbers)
    fmt.Println(result)

    sort.Ints(numbers)
    fmt.Println(numbers)

    result = sort.IntsAreSorted(numbers)
    fmt.Println(result)
}