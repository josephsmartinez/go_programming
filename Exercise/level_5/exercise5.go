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

}
