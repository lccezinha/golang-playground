package main

import "fmt"

func main() {
	var real float32
	const dolar float32 = 3.8

	fmt.Println("Informe o valor em reais R$: ")
	fmt.Scanf("%f", &real)

	cotation := real / dolar

	fmt.Printf("O valor de R$ %.2f em doláres é U$ %.2f\n", real, cotation)
}