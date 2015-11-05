package main 

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		switch i {
			case 0: fmt.Println(0, "Zero")
			case 1: fmt.Println(1, "Um")
			case 2: fmt.Println(2, "Dois")
			case 3: fmt.Println(3, "Três")
		}
	}
}