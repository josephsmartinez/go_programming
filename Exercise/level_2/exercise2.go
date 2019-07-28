package main

import (
	"fmt"
)

var palindrome string = "tacocat"
var not_palindrome string = "palindrom"

func main() {
	fmt.Println("Comparing words:", palindrome, "and", not_palindrome)
	isequal := bool(palindrome == not_palindrome)
	fmt.Println("isequal: ", isequal)
	lessthan_equalto := bool(palindrome <= not_palindrome)
	fmt.Println("lessthan_equalto: ", lessthan_equalto)
	greaterthan_equalto := bool(palindrome >= not_palindrome)
	fmt.Println("greaterthan_equalto: ", greaterthan_equalto)
	notequal := bool(palindrome != not_palindrome)
	fmt.Println("notequal: ", notequal)
	greaterthan := bool(palindrome < not_palindrome)
	fmt.Println("greaterthan: ", greaterthan)
	lessthan := bool(palindrome > not_palindrome)
	fmt.Println("lessthan: ", lessthan)

}
