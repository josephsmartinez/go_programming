package main

import (
	"fmt"
)

var x int = 12

func main() {
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Hexidecimal: %X\n", x)
	fmt.Printf("Binary: %b\n", x)
	// Bit shifting
	x = x << 1
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Hexidecimal: %X\n", x)
	fmt.Printf("Binary: %b\n", x)

	x = x >> 4
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Hexidecimal: %X\n", x)
	fmt.Printf("Binary: %b\n", x)
	/*
		Go apparently doesn't accept the 0b notation for binary integers. I was just using it for the example. In decimal, 8 >> 1 is 4, and 8 << 1 is 16. Shifting left by one is the same as multiplication by 2, and shifting right by one is the same as dividing by two, discarding any remainder.
	*/
}
