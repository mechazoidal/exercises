package main

import ("math";
"fmt";
)

func Sqrt(x float64) float64 {
  z := float64(1)
  for i := 0; i < 10; i++ {
    z = (z - (( z*z - x ) / z*2) )
  }
  return z
}

func main() {
  fmt.Println(Sqrt(4))
  fmt.Println(mainmath.Sqrt(4))
}
