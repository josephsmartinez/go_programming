package main

import (
	"fmt"
)

type person struct {
	firstname, lastname string
	ice_cream_list      []string
}

type worker struct {
	person
	job    string
	salary int
}

func main() {

	p1 := person{
		firstname:      "Mike",
		lastname:       "Gummpy",
		ice_cream_list: []string{"Coco", "Vanila", "Stawberry"},
	}

	p2 := worker{
		person: person{
			firstname:      "John",
			lastname:       "Doe",
			ice_cream_list: []string{"Mango", "Rum", "Random"},
		},
		job:    "ice cream maker",
		salary: 250000,
	}
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)

	for k, v := range p1.ice_cream_list {
		fmt.Printf("%v %v\n", k+1, v)

	}

	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	fmt.Println(p2.ice_cream_list)
	fmt.Println("\n\n")

	/*
		Take the code from the previous exercise, then
		store the values of type person in a map with the
		key of last name. Access each value in the map.
		Print out the values, ranging over the slice.
	*/

	p3 := person{
		firstname:      "Jen",
		lastname:       "Stepperfind",
		ice_cream_list: []string{"Time", "Space", "Now"},
	}

	p4 := person{
		firstname:      "Humberto",
		lastname:       "Stevens",
		ice_cream_list: []string{"Here", "There", "Over"},
	}

	m := map[string]person{
		p3.firstname: p3,
		p4.lastname:  p4,
	}

	for _, v := range m {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for i, val := range v.ice_cream_list {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
