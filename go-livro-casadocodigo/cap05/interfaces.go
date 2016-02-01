package main

import (
  "fmt"
  "time"
)

type Operation interface{
  Calculate() int
}

type Sum struct {
  n1, n2 int
}

type Subtration struct {
  n1, n2 int
}

type Age struct {
  yearBirthday int
}

// Type Sum, implementa a interface Operation implicitamente, apenas implementando seu método Calculate()
func (s Sum) Calculate() int {
  return s.n1 + s.n2
}

func (s Subtration) Calculate() int {
  return s.n1 - s.n2
}

func (a Age) Calculate() int {
  return time.Now().Year() - a.yearBirthday
}

func (s Sum) String() string {
  return fmt.Sprintf("%d + %d", s.n1, s.n2)
}

func (s Subtration) String() string {
  return fmt.Sprintf("%d - %d", s.n1, s.n2)
}

func (s Age) String() string {
 return fmt.Sprintf("Age since of: %d", s.yearBirthday)
}

func sum(opps []Operation) int {
  accumulate := 0

  for _, op := range opps {
    value := op.Calculate()
    fmt.Printf("%v = %d\n", op, value)
    accumulate += value
  }

  return accumulate
}

func main() {
  // Assinatura do método Calculate do tipo Sum, satisfaz a condição da interface Operation
  // desse modo, Sum é considerado uma Operation.
  // assim, pode-se criar um variável do tipo Open e executar os métodos de Sum.

  // var sum Operation
  // sum = Sum{10, 15}

  // fmt.Printf("%v = %d\n", sum, sum.Calculate())

  opp := make([]Operation, 4)
  opp[0] = Sum{10, 15}
  opp[1] = Subtration{20, 10}
  opp[2] = Subtration{30, 15}
  opp[3] = Sum{100, 5}

  fmt.Printf("Value %v \n", sum(opp))

  ages := make([]Operation, 3)
  ages[0] = Age{1990}
  ages[1] = Age{1964}
  ages[2] = Age{1923}

  fmt.Printf("Value %v \n", sum(ages))
}