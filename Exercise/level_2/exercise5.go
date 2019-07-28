package main

import (
	"fmt"
)

var word string = "Hello World!"
var raw_string string = `Hello \n \t word raw string`

func main() {
	for i := 0; i < len(word); i++ {
		fmt.Printf("%q %U\n", word[i], word[i])
	}
	fmt.Println(raw_string)
}
