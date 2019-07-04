package main

import "fmt"

func main() {

	// Simple Map
	a := map[string]string{
		"foo": "bar",
	}
	fmt.Println(a)

	// Initialize a map
	routes := make(map[string]int)
	routes["route66"] = 66
	route66 := routes["route66"]
	fmt.Println(route66)

	// Map builtin functions
	countries := map[int]string{
		0: "Africa",
		1: "America",
		2: "Canada",
		3: "France",
		4: "Russia",
		5: "Chaina",
	}

	// len()
	length := len(countries)
	// delete()
	delete(countries, 1)
	newLength := len(countries)
	// Two value assignment
	// If that key doesn't exist, i is the value type's zero value (0)
	element, isAmericaInMap := countries[1]
	_, isRussiaInMap := countries[4]
	fmt.Printf("America in map? %v\nRussia in map? %v\nLentgh: %v\nLength after deleting America: %v\nReturn value after deleting America: %v\n", isAmericaInMap, isRussiaInMap, length, newLength, element)

}
