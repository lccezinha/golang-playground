package main 

import "fmt"

type User struct {
    Id int
    Name, Location string
}

type Player struct {
    User
    GameId int
}

func main() {
    p := Player{}
    p.Id = 1
    p.Name = "Xunda"
    p.Location = "Brasil"
    p.GameId = 2

    fmt.Printf("%+v", p)
}