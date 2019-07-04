package main

import "fmt"

// Standard function with return values
// NOTE: If consecutive parameters are of the same type, omit re-defining type
func calculateBill(price, tip float32) float32 {
	var totalPrice = price / (1 - tip)
	return totalPrice
}

// Multiple return values

// Variadic parameter
func permIndependantStages(stages int, choices ...int) int {
	var product = choices[0]
	if stages == 0 {
		return 0
	} else if stages == 1 {
		return choices[0]
	} else {
		for i := 0; i < stages-1; i++ {
			product *= choices[i+1]
		}
	}
	return product
}

func main() {

	var price, tip float32 = 20.00, .20
	bill := calculateBill(price, tip)
	fmt.Println("Total Bill is: ", bill)

	stages := 4
	numOptions := permIndependantStages(stages, 2, 3, 5, 3)
	fmt.Println(numOptions)

}
