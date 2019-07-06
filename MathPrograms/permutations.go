package main

import "fmt"

func permIndependantStages(choices ...int) int {
	var product int
	if len(choices) == 0 {
		return 0
	} else if len(choices) == 1 {
		return choices[0]
	} else {
		product = choices[0]
		for i := 0; i < len(choices)-1; i++ {
			product *= choices[i+1]
		}
	}
	return product
}

func main() {

	x := permIndependantStages(2, 3, 5, 3)
	fmt.Println(x)
}
