package main

import "fmt"

func spliter(numbers []int, odd, even chan<- int, ready chan<- bool) {
  for _, num := range numbers {
    if num % 2 == 0 {
      even <- num
    } else {
      odd <- num
    }
  }

  ready <- true
}

func main() {
  odd, even := make(chan int), make(chan int)
  ready := make(chan bool)
  numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  go spliter(numbers, odd, even, ready)

  var resultOdd, resultEven []int
  end := false

  for !end {
    select {
      case n := <-odd:
        resultOdd = append(resultOdd, n)
      case n := <-even:
        resultEven = append(resultEven, n)
      case end = <-ready:
    }
  }

  fmt.Printf("Odd: %v | Even: %v\n", resultOdd, resultEven)
}