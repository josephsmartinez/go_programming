// every package start with a main package
package main

// import functions
import "fmt"

// main method
func main() {
	fmt.Println("Hello World!!!")
	foo()

	// simple for loop the print even numbers
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
