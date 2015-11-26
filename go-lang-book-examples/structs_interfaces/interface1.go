// Method

// (c *Circle) é o receiver do método
// area() é o método
// fazendo dessa forma é possível fazer c.area() ~ como se fosse um objeto

/*
func (c *Circle) area() float64 {
  return c.x * c.x
}
*/

package main

import (
  "fmt"
  "math"
)

type Shaper interface {
  area() float64
}

type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}
 
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

// criando um método com receive, não é preciso passar como ponteiro
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func totalAreas(shapes ...Shaper) float64 {
  var area float64

  for _, shape := range shapes {
    area += shape.area()
  }

  return area
}

func main() {
  r := Rectangle{0, 0, 10, 10}
  c := Circle{0, 0, 5}

  fmt.Println(r.area())
  fmt.Println(c.area())
  fmt.Println(totalAreas(&c, &r))
}