package main

import "fmt"

func main() {
  a := make([]int, 5)
  printSlice("a -> ", a)

  b := make([]int, 5, 10)
  printSlice("b ->", b)
}

func printSlice(s string, x []int){
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
