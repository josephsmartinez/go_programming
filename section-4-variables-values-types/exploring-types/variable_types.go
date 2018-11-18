package main

import "fmt"

var y = 42
var message_1 = "Shaken, not stirred"
var message_2 string = "Assigned the value of TYPE string"
var message_3 = `James said, "I'll take the Drink now good sir"`

func main() {
	fmt.Println(y)

	// To print the type of a value, use %T.
	fmt.Println("%T", message_1)
	fmt.Println(message_2)
	fmt.Println(message_3)
}
