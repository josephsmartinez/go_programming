package main

import "fmt"

func main() {
	True := true
	False := false

	if True && True {
		fmt.Printf("%v", True)
	} else if True && False {
		fmt.Printf("%v", False)
	}
}
