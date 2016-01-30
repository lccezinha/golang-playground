package main

import (
  "fmt"
  "sort"
)

func simpleMap() {
  numbers := make(map[string]int)

  numbers["one"] = 1
  numbers["two"] = 2

  fmt.Println(numbers, len(numbers), numbers["one"], numbers["two"])
}

func simpleMapAlt() {
  numbers := map[int]string{}

  numbers[1] = "One"
  numbers[2] = "Two"

  fmt.Println(numbers, len(numbers), numbers[1], numbers[2])
}

func simpleMapAltTwo() {
  numbers := make(map[int]string, 10)

  numbers[1] = "One"
  numbers[2] = "Two"

  fmt.Println(numbers, len(numbers), numbers[1], numbers[2])
}

type State struct {
  Name, Capital string
  Population int
}

func states() {
  states := make(map[string]State, 3)

  states["PR"] = State{"Paraná", "Curitiba", 1000}
  states["SC"] = State{"Santa Catarina", "Florianópolis", 2000}
  states["RS"] = State{"Rio Grande do Sul", "Porto Alegre", 2500}
}

func findKeys() {
  states := make(map[string]State, 3)

  states["PR"] = State{"Paraná", "Curitiba", 1000}
  states["SC"] = State{"Santa Catarina", "Florianópolis", 2000}
  states["RS"] = State{"Rio Grande do Sul", "Porto Alegre", 2500}

  parana := states["PR"]
  fmt.Println(parana)
}

func handleInvalidKeys() {
  states := make(map[string]State, 3)

  states["PR"] = State{"Paraná", "Curitiba", 1000}
  states["SC"] = State{"Santa Catarina", "Florianópolis", 2000}
  states["RS"] = State{"Rio Grande do Sul", "Porto Alegre", 2500}

  fmt.Println(states["SP"]) // retorna o Zero value -> {  0}

  // Pode verificar antes

  sp, ok := states["SP"]
  if ok {
    fmt.Println(sp)
  } else {
    fmt.Println("SP not found")
  }
}

func deleteItem() {
  states := make(map[string]State, 3)

  states["PR"] = State{"Paraná", "Curitiba", 1000}
  states["SC"] = State{"Santa Catarina", "Florianópolis", 2000}
  states["RS"] = State{"Rio Grande do Sul", "Porto Alegre", 2500}

  fmt.Println(states)

  delete(states, "PR")

  fmt.Println(states)

  delete(states, "SP")

  fmt.Println(states)
}

func iterator() {
  states := make(map[string]State, 3)

  states["PR"] = State{"Paraná", "Curitiba", 1000}
  states["SC"] = State{"Santa Catarina", "Florianópolis", 2000}
  states["RS"] = State{"Rio Grande do Sul", "Porto Alegre", 2500}

  for _, state := range states {
    fmt.Printf("%s (%s) has %d n00bs \n", state.Name, state.Capital, state.Population)
  }
}

func order() {
  sq := make(map[int]int, 15)

  for n := 1; n <= len(sq); n++ {
    sq[n] = n * n
  }

  numbers := make([]int, 0, len(sq)) // tamanho inicial: 0, capacidade: 15

  for n := range sq {
    numbers = append(numbers, n)
  }
  sort.Ints(numbers)

  for _, number := range numbers {
    q := sq[number]
    fmt.Printf("%d ^ 2 = %d\n", number, q)
  }
}

func main() {
  simpleMap()
  simpleMapAlt()
  simpleMapAltTwo()
  states()
  findKeys()
  handleInvalidKeys()
  deleteItem()
  iterator()
  order()
}