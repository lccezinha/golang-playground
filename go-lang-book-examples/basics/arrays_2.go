package main 

import "fmt"

// func main() {
// 	var x [5]float64
// 	total := 0.0

// 	x[0] = 10.0
// 	x[1] = 92.5
// 	x[2] = 39.2
// 	x[3] = 94.1
// 	x[4] = 88.2

// 	for i := 0; i < 5; i++ {
// 		total += x[i]
// 	}

// 	fmt.Println(total / 5)
// }

// func main() {
// 	var x [5]float64
// 	total := 0.0

// 	x[0] = 10.0
// 	x[1] = 92.5
// 	x[2] = 39.2
// 	x[3] = 94.1
// 	x[4] = 88.2

// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}

// 	// fmt.Println(total / len(x)) // => falha pq total é float64 e len(x) é int, precisa converter
// 	fmt.Println(total / float64(len(x)))
// }

func main() {
	x := [5]float64{10.0, 92.5, 39.2, 94.1, 88.2}
	total := 0.0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
}
