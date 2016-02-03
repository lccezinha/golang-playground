package main

import (
  "fmt"
  "regexp"
  "strings"
)

func main() {
  expr := regexp.MustCompile("\\b\\w")

  transformer := func(str string) string {
    return strings.ToUpper(str)
  }

  text := "luiz cezer marrone filho"

  fmt.Println(transformer(text))
  fmt.Println(expr.ReplaceAllStringFunc(text, transformer))
}