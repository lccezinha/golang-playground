package main

import "fmt"

func addItem() {
  s := make([]int, 0)
  s = append(s, 10)

  fmt.Println(s)
}

func addItemToStart() {
  s := []int{11, 12, 13}
  n := []int{10}

  n = append(n, s...)

  fmt.Println(n)
}

func addItemToStartTwo() {
  s := []int{1, 2, 3}
  s = append([]int{15}, s...)

  fmt.Println(s)
}

func addItemToMiddle() {
  s := []int{1, 2, 3, 6, 7}
  v := []int{4, 5}
  s = append(s[:3], append(v, s[3:]...)...)

  fmt.Println(s)
}

func removeItemFromStart() {
  s := []int{1, 2, 3, 4, 5}

  fmt.Println(s[1:])
}

func removeItemFromEnd() {
  s := []int{1, 2, 3, 4, 5}

  fmt.Println(s[:4])
}

func removeItemFromMiddle() {
  s := []int{1, 2, 3, 4, 5}
  v := append(s[:2], s[4:]...)

  fmt.Println(v)
}

func double() {
  numbers := []int{1, 2, 3, 4, 5}
  double := make([]int, len(numbers))

  copy(double, numbers)

  for i := range double {
    double[i] *= 2
  }

  fmt.Println(numbers, double)
}

func main() {
  addItem()
  addItemToStart()
  addItemToStartTwo()
  addItemToMiddle()
  removeItemFromStart()
  removeItemFromEnd()
  removeItemFromMiddle()
  double()
}