package main

import "fmt"

// Standard function
func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

// Standard function with return values
// NOTE: If consecutive parameters are of the same type, omit re-defining type
func calculateBill(price, tip float32) float32 {
	var totalPrice = price / (1 - tip)
	return totalPrice
}

// Multiple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// Variadic parameter
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

// Returning a function from a function
func one() func() int {
	return func() int {
		return 451
	}
}

// Callbacks and functional programming (NOT RECOMMENDED)
func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

// Recursion

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

////////////////////////////
// MAIN FUNCTION FOR TESTING
////////////////////////////
func main() {

	var price, tip float32 = 20.00, .20
	bill := calculateBill(price, tip)
	fmt.Println("Total Bill is: ", bill)

	numOptions := permIndependantStages(2, 3, 5, 3)
	fmt.Println(numOptions)

	// Unfurling a slice
	options := []int{2, 3, 5, 3}
	numberOfOptions := permIndependantStages(options...)
	fmt.Println(numberOfOptions)

	// Defer
	defer foo()
	bar()

	// Anonymous Functions
	func() {
		fmt.Println("anonymous function")
	}()

	func(x int) {
		fmt.Println("anonymous function() = ", x)
	}(10)

	// Function expressions
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)

	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	n := factorial(4)
	fmt.Println(n)

}
