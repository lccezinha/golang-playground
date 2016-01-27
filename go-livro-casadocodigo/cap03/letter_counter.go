package main

import (
  "fmt"
  "os"
  "strings"
)

func buildStats(words []string) map[string]int {
  results := make(map[string]int)

  for _, word := range words {
    firstLetter := strings.ToUpper(string(word[0]))
    results[firstLetter] += 1
  }

  return results
}

func printResults(results map[string]int) {
  fmt.Println("First letter word counter: \n")

  for key, value := range results {
    fmt.Printf("%s = %d\n", key, value)
  }
}

func main() {
  if len(os.Args) < 1 {
    fmt.Println("no args")
    os.Exit(1)
  }

  words := os.Args[1:]
  statistcs := buildStats(words)
  printResults(statistcs)
}

