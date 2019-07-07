package main

import "fmt"

func main() {
	// Slice Composite Literal
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// Slice Make
	b := make([]string, 3, 3)
	fmt.Println(b)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Apppend to Slice
	b = append(b, "Value1", "Value1", "Value1", "Value1", "Value1")

	// 2 Demensional Slice
	c := []float32{1.0, 2.0, 3.0, 4.0}
	d := []float32{5.0. 6.0, 7.0, 8.0}
	e := [][]float32{c,d}
	
	// 2 Demensional Slice Matrix
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

// Slicing a slice
  f := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// Declaring Empty Slices
	var t []string  // empty slice
	t := []string{} // nil slice


// Slice methods
// len()
// cap()
// append()
// copy()
	
}
