package main

import "fmt"

func main() {
	// func Println(a ...interface{}) (n int, err error)
	fmt.Println("Hello", true, 100)

	// catch the return information from the "fmt" method
	n, err := fmt.Println("Hello", true, 100)
	fmt.Println(n)
	fmt.Println(err)

	// underscores will throw return errors into null
	m, _ := fmt.Println("Hello", true, 100)
	fmt.Println(m)
}
