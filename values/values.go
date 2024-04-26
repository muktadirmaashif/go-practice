package main

import (
  "fmt"
  // "math"
)

const s string = "constant"
// in case of var, can be declared first and then assigned later, but with const, needs to be done together.

func main() {
  fmt.Println(s)


  // one liner
  for j:=1; j<=11; j++ {
    if j%2 == 0 { 
      continue 
    }
    fmt.Print(j, " ")
    if j == 10 {
      break 
    }
  }
  fmt.Println()

  

  // const d = 45
  // fmt.Println(math.Sin(d))
}
