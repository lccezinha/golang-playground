package main

import "fmt"

type Shaper interface {
  Area() int
}

type Rectangle struct {
   length, width int
}

func (r Rectangle) Area() int {
   return r.length * r.width
}

type Square struct {
  side int
}

func (sq Square) Area() int {
  return sq.side * sq.side
}

func main() {
  r := Rectangle{length:5, width:3} //define a new Rectangle instance with values for its properties
  fmt.Println("Rectangle details are: ",r)
  fmt.Println("Rectangle's area is: ", r.Area())

  var s Shaper
  s = r // igual fazer ~ s = Shaper(r) ~ desde que Go identifica que 'r' encontra a interface Shaper pelo fato de ela ter o mesmo método que a interface
  fmt.Println("Area os shaper(Rectangle) is: ", s.Area())

  q := Square{side: 3}
  fmt.Println("Area os shaper is: ", q.Area())

  s = q // igual fazer ~ s = Shaper(r) ~ desde que Go identifica que 's' encontra a interface Shaper pelo fato de ela ter o mesmo método que a interface
  fmt.Println("Area os shaper is (Square): ", s.Area())
}
