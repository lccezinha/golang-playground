package main

import "fmt"

type User struct {
    FirstName, LastName string
}

type Customer struct {
    Id int 
    Fullname string
}

func (u *User) Name() string {
    return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (c *Customer) Name() string {
    return c.Fullname
}

type Namer interface {
    Name() string
}

func Greet(n Namer) string {
    return fmt.Sprintf("Olá %s", n.Name())
}

func main() {
    u := &User{FirstName: "Cezer", LastName: "Luiz"}
    fmt.Println(Greet(u))

    c := &Customer{1, "Cezinha"}
    fmt.Println(Greet(c))
}