package main

import (
	"fmt"
	"runtime"
)

func main() {

	// traditional for loops
	fmt.Println("for loop -------")
	sum := 0
	for i := 0; i < 100; i++ {
		sum += 2
	}
	fmt.Printf("The sum is: %v\n", sum)

	// range for loop
	fmt.Println("range loop -------")
	person := map[string]string{
		"firstName": "James",
		"lastName":  "Bond",
		"city":      "London",
	}
	for key, val := range person {
		fmt.Printf("%v %v\n", key, val)
	}

	// while
	fmt.Println("while loop -------")
	evens := 0
	for evens < 100 {
		if evens%2 == 0 {
			evens += evens
		}
		evens++
	}
	fmt.Printf("%v\n", evens)

	// do while
	fmt.Println("do while loop -------")
	twice := 2
	for {
		twice--
		fmt.Printf("%v\n", twice)
		if twice < 1 {
			break
		}
	}

	// switch
	fmt.Println("switch statement -------")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("os x\n")
	case "linux":
		fmt.Printf("linux\n")
	case "windows":
		fmt.Printf("window\n")
	default:
		fmt.Printf("other\n")
	}

	fmt.Println("switch statement -------")
	pet := "dog"
	switch pet {
	case "cat":
		fmt.Printf("cat\n")
	case "monkey":
		fmt.Printf("monkey\n")
	case "dog":
		fmt.Printf("dog\n")
	}

	switch {
	case true:
		fmt.Println("switch with no condition")
	}

	// if, if else, else
	fmt.Println("if, if else, else -------")
	if 10%2 == 0 {
		fmt.Println("works 1")
	} else if true {
		fmt.Print("works 2")
	} else {
		fmt.Println("works 3")
	}

	// if with a short statement
	fmt.Println("short if statement -------")
	if v := true; v == true {
		fmt.Printf("To be or not to be? %v\n", v)

	} else if v := false; v == true {
		fmt.Printf("To be or not to be? %v\n", v)
	}

	// defer statement
	defer fmt.Println("2 was defered")
	fmt.Println("1")

}
