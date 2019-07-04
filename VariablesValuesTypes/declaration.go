package main

import "fmt"

func main() {
	// assigning a variable
	x := 100
	fmt.Println(x)
	// reassignment
	x = 50
	fmt.Println(x)

	// assigning a string values
	y := "Hello"
	print(y)

}

// Function prints a string
func printstring(x string) {
	fmt.Println(x)
}

// Function print a number
func printint(x int) {
	fmt.Println(x)
}
