package main

import "fmt"

type User struct {
    Id             int
    Name, Location string
}

func (u *User) Greetings() string {
    return fmt.Sprintf("Hi %s from %s. My id is %d \n", u.Name, u.Location, u.Id)
}

type Player struct {
    *User
    GameId int
}

func main() {
    p := &Player{
        User: &User{Id: 1, Name: "Xunda", Location: "Brasil"},
        GameId: 2,
    }

    fmt.Printf("%+v", p.Greetings())
}