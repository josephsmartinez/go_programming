package main

import "fmt"

type person struct {
	name string
}

var people []person

func main() {
	person1 := person{name: "mike"}
	fmt.Println(person1.name)

	// person2 := person{name: "sassy"}
	// person3 := person{name: "tim"}
	// person4 := person{name: "walter"}
	// person5 := person{name: "kim"}

	people = append(people, person1)

	fmt.Println(people[0])

}
