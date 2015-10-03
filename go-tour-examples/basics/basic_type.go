// Tipos básicos

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // pseudônimo para uint8

// rune // pseudônimo para int32
//      // representa um ponto de código Unicode

// float32 float64

// complex64 complex128

/////////////////////////////////////

// Variáveis declaradas sem um valor inicial explicitado darão seu valor zero.

// O valor zero é:

//   0 para tipos numéricos,
//   false para tipos boleanos, e
//   "" (string vazia) para strings.

package main

import "fmt"

func main(){
  var i int
  var f float32
  var b bool
  var s string

  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
