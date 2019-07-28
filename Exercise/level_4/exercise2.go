package main

import (
	"fmt"
)

func main() {
	xSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, v := range xSlice {
		fmt.Printf("Type: %T Value: %v\n ", v, v)
	}

	ySlice := make([]int, 10, 10)
	printSlice("x", ySlice)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
