package main

import "fmt"

func foo(x ...int) int {
	var sum int
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := foo(numbers...)
	fmt.Println(sum)
}
