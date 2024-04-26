package main 

import (
  "fmt"
  "time"
)


func main() {
  i:=2
  // Print = without any space and new line. needs to be mentioned them explicitly. Unlike Println
  fmt.Print("Write ",i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Friday:
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("It's a weekday.")
  }

  t := time.Now()
  fmt.Print("Time is: ", t.Weekday(), ", ", t.Hour()-12,":",t.Minute(), " PM","\n")
  
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default: 
    fmt.Println("It's after noon")
  }
}
