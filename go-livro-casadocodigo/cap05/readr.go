package main

import (
  "fmt"
  "io"
)

type Readr struct {}

func (r Readr) Read(p []byte) (int, error) {
  p[0] = 'B'
  p[1] = 'A'
  p[2] = 'T'
  p[3] = 'A'
  p[4] = 'T'
  p[5] = 'A'

  return len(p), nil
}

func readString(r io.Reader) string {
  p := make([]byte, 6)
  r.Read(p)

  return string(p)
}

func main() {
  reader := Readr{}

  fmt.Println(readString(reader))
}