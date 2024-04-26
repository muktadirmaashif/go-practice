package main
import (
  "fmt"
)
func main() {
  var arr [4]int    // only declare, no initialization
  fmt.Println("arr: ",arr)

  arr[3] = 100
  fmt.Println("arr: ",arr)
  fmt.Println("in 4th: ",arr[3])
  fmt.Println("length: ",len(arr))

  b := [5]int{1,2,3,4,5}
  fmt.Println("b: ",b)

  var tD [2][3]int
  for i:=0; i<2; i++ {
    for j:=0; j<3; j++ {
      tD[i][j] = i+j 
    }
  }
  fmt.Println("two Dimensinal array: ", tD)
  // with Println, displayed as [v1 v2 v3] form
  


}
