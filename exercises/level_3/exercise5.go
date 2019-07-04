package main

import "fmt"

var (
	start int = 10
	end   int = 100
)

func main() {
	mod4 := 0
	for {
		mod4 = start % 4
		fmt.Printf("Remainder is:\n%v/4=%v\n", start, mod4)
		start++
		if start > end {
			break
		}
	}

}
