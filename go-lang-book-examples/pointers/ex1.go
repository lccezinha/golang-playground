package main 

import "fmt"

func swap(pointerToX, pointerToY *int) {
	v := *pointerToX
	*pointerToX = *pointerToY
	*pointerToY = v
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}