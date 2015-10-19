package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (p Person) sayName() string{
  if p.Name == "" {
    return "Nome nulo"
  } else {
    return p.Name
  }
}

func (p Person) isAdult() bool {
  if p.Age >= 18 {
    return true
  } else {
    return false
  }
}

func (p Person) message() string {
  if p.isAdult() {
    return "Adulto"
  } else {
    return "Juvenil"
  }

}

func main(){
  person := Person{"Luiz", 11}

  fmt.Println(person.Name)
  fmt.Println(person.Age)
  fmt.Println(person.message())
}
