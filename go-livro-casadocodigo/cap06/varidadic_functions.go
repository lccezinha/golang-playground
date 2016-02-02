package main

import (
  "fmt"
  "os"
)

func createFiles(baseDrir string, files ...string) {
  for _, file := range files {
    filePath := fmt.Sprintf("%s/%s.%s", baseDrir, file, "txt")

    arq, err := os.Create(filePath)
    defer arq.Close()

    if err != nil {
      fmt.Printf("Erro to create file : %s - %v \n", file, err)
      os.Exit(1)
    }

    fmt.Printf("File %s created. \n", arq.Name())
  }
}

func main() {
  tmp := os.TempDir()

  createFiles(tmp)
  createFiles(tmp, "file1")
  createFiles(tmp, "file2", "file3")
}