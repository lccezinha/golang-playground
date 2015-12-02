package main

import "os"
import "errors"

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		errors.New("Não pode abrir o arquivo.")
		return 
	}

	defer file.Close()

	file.WriteString("teste maroto")
}