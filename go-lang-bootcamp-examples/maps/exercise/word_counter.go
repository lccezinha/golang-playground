package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    words := make(map[string]int)
    splitedWords := strings.Fields(s)

    for _, word := range splitedWords {
        words[word] = len(word)
    }

    return words
}

func main() {
    text := "xunda maneiro"
    result := WordCount(text)

    fmt.Println(result)
}