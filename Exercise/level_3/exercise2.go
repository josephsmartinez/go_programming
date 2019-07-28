package main

import "fmt"

const A = 65
const Z = 90

func main() {

	letter := A
	for {
		for i := 3; i >= 0; i-- {
			fmt.Printf("%U %c\n", letter, letter)
		}
		letter++
		if letter > Z {
			break
		}
	}
}
