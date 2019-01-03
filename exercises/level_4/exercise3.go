package main
import (
  "fmt"
)
func main() {
  matrix := [][]int{
    {1,2,3,4,5,6,7,8,9},
    {1,3,5,7,9},
    {2,4,6,8},
    {1,3,5,7},
  }

for _, i := range matrix {
  for _, j := range i {
    fmt.Printf("%v ", j)
  }
  fmt.Println("")
}

}
