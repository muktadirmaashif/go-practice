package main
import (
  "fmt"
  //"slices"
)
func main() {
 var a []string // Declaration is similar to array. diff - [] is empty.
 // fmt.Println("slice: ",a,a==nil,len(a)==0)
 // empty slice. A slice contains three elements. Pointer to the underlying array, 
 // length-number of elements of the array, capacity-max number of elements the array can have. 
 a = make([]string, 3, 5) // slice initialization
 fmt.Println("slice: ", a, "len:", len(a), "cap:", cap(a))
 var s []string
 s = append(s, "d", "e", "f", "g", "h", "i")
 fmt.Println("append: ",s, "length: ",len(s), "cap: ", cap(s))

 copy(a,s) // copy(dest,src)
 fmt.Println("copy:",s)

 fmt.Println("slice: ")
 l := s[1:4] // slicing a slice. I know, contradictory. s[from this : untill this]
 fmt.Println("l is s[1:4]:",l)

 l = s[:4] 
 fmt.Println("l is s[:4]:",l)

 l = s[1:] 
 fmt.Println("l is s[1:]:",l)
 l = s[:] 
 fmt.Println("l is s[:] (similar to copy):",l)



 
 
}
