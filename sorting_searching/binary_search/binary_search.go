package main

import "fmt"

// Function return a boolen value when number is found.
func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
  // if the list is empty OR the first element is NOT the equal to the needle.
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true

}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
