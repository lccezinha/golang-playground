package main

import (
	"fmt"
	"os"
)

func main() {
	file, err	:= os.Create("xunda.txt")
	if err != nil { return }

  file, err = os.Open("xunda.txt")
  if err != nil {
    return
  }
  defer file.Close()

  stat, err := file.Stat()
  if err != nil {
    return
  }

  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)
  if err != nil {
    return
  }

  str := string(bs)
  fmt.Println(str)
}