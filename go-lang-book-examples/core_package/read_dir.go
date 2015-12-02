package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Lista só os diretórios dentro do dir /home/cezinha/Dev

	// dir, err := os.Open("/home/cezinha/Dev/")
	// if err != nil { return }
	// defer dir.Close()

	// fileInfos, err := dir.Readdir(-1)
	// if err != nil { return }

	// for _, file := range fileInfos {
	// 	fmt.Println(file.Name())
	// }

	// Passa por todas pastas e subpastas
	filepath.Walk("/home/cezinha/Dev", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}