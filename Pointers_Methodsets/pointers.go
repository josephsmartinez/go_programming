package main

import (
	"fmt"
)

func main() {
	// Declaring pointers
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// Zero value of a pointer
	{
		a := 25
		var b *int
		if b == nil {
			fmt.Println("b is", b)
			b = &a
			fmt.Println("b after initialization is", b)
		}
	}

	// Creating pointers using the `new`function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	// Dereferencing a pointer
	{
		b := 255
		a := &b
		fmt.Println("address of b is", a)
		fmt.Println("value of b is", *a)
	}

	// Passing pointer to a function
	{
		a := 58
		fmt.Println("value of a before function call is", a)
		b := &a

		func(val *int) {
			*val = 55
		}(b)

		fmt.Println("value of a after function call is", a)
	}

	// Pointer/Non-Pointer Receivers and Pointer Values

	word := "Hello World!"
	fmt.Printf("%T , %v, %v\n", word, word, &word)

	// Non-Pointer Receivers
	func(word string) {
		fmt.Printf("%T , %v, %v\n", word, word, &word)
	}(word)

	// Pointer Receivers
	func(word *string) {
		fmt.Printf("%T , %v, %v\n", word, &word, *word)
	}(&word)

}
