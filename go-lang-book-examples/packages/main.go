package main 

import "fmt"
import "./mathme"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := mathme.Average(xs)

	fmt.Println(avg)
}