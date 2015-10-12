package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main(){
  m = make(map[string]Vertex)
  m["positions"] = Vertex{123.52, -321.321}

  fmt.Println(m["positions"])
}
