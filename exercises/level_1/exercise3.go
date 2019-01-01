package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	statement := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	statement2 := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(statement)
	fmt.Println(statement2)
}
