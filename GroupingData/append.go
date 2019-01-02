package main

import (
	"fmt"
)

func main() {
	y := []int{1, 2, 3}
	x := []int{4, 5, 6, 7, 8, 9}
	x = append(y, x...)
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		if x[i]%2 == 0 {
			// Remove even number
			fmt.Println(x[i])
			x = append(x[:i], x[i+1:]...)
		}
	}
	fmt.Println(x)
}
