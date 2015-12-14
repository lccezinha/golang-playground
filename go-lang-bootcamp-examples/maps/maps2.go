package main

import "fmt"

type Vertex struct {
    Long, Lat float64
}

var m map[string]Vertex

func main() {
    m := make(map[string]Vertex)
    m["Xunda"] = Vertex{Long: 12.1, Lat: 311.2}
    fmt.Println(m["Xunda"], m)    
}