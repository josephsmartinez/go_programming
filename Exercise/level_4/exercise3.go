package main

import (
	"fmt"
)

func main() {

	x := []int{}

	for i := 42; i <= 51; i++ {
		x = append(x, i)
	}
	print(x)
	print(x[:5])
	print(x[5:])
	print(x[2:7])
	print(x[1:6])
	print(x)

}
func print(x []int) {
	fmt.Println(x)
}
