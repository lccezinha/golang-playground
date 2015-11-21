package main

import "fmt"


func main() {
	numbers := []float64{10.0, 92.5, 39.2, 94.1, 88.2}
	fmt.Println(average(numbers))
}

func average(numbers []float64) float64{
	total := 0.0
	for _, value := range numbers {
		total += value
	}
	average := total / float64(len(numbers))
	return average
}
