package main

import "fmt"

var z = "This variable can be delared outside of the method."
var un_assigned int

func main() {
	// DECLARE a varaible and ASSIGN a VALUE
	x := 42
	fmt.Println(x)
	var y = 43
	fmt.Println(y)

	// print the global variable
	foo()
}

func foo() {
	fmt.Println(z)
}
