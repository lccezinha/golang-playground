package main

import "fmt"

//panic() irá forçar um run time error
//recover() manipula o erro lançado por panic()
//recover() irá para panic() e retornar o valor passado a chamada de panic()

func main() {
	// Porém dessa forma não irá funcionar, por que panic() irá parar a execução imediatamente
	// panic("HELP")
	// message := recover()
	// fmt.Println(message)

	// Dessa forma irá forçar que recover execute
	defer func() {
		message := recover()
		fmt.Println(message)	
	}()
	panic("HELP")
}

