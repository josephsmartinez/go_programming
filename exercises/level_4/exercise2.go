package main
import (
  "fmt"
)
func main() {
   t := []string{"a", "b", "c", "d", "e", "f", "g"}
   for _, v := range t {
     fmt.Printf("Type: %T Value: %v\n ", v, v)
   }
}
