package main
import (
  "fmt"
)
func main() {
  array := [5]string{"Hello", "World", "Welcome", "to", "Go"}
  for k, v := range array{
    fmt.Printf("Index: %v Type: %T String Value: %v\n ", k, v, v)
  }
}
