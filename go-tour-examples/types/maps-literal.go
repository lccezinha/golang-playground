// package main

// import "fmt"

// type Vertex struct {
//   Lat, Long float64
// }

// var m = map[string]Vertex{
//   "positions A": Vertex{123.52, -321.321},
//   "positions B": Vertex{3412.21, 32.2212},
// }

// func main(){
//   fmt.Println(m)
// }


package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "positions A": {123.52, -321.321},
  "positions B": {3412.21, 32.2212},
}

func main(){
  fmt.Println(m)
}
