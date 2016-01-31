package main

import "fmt"

type MyFile struct {
  name string
  size float32
  lines int
}

func main() {
  file := &MyFile{name: "file.txt", size: 32, lines: 1500}
  fileTwo := &MyFile{size: 150}
  fileThree := &MyFile{}
  fileFour := &MyFile{"some.go", 200, 2500}

  fmt.Println(file)
  fmt.Println(fileTwo)
  fmt.Println(fileThree)
  fmt.Println(fileFour)
}