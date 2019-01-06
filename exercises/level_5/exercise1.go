package main

import (
	"fmt"
)

type person struct {
	firstname, lastname string
	ice_cream_list      []string
}

func main() {

	p1 := person{
		firstname:      "Mike",
		lastname:       "Gummpy",
		ice_cream_list: []string{"Coco", "Vanila", "Stawberry"},
	}
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)

	for k, v := range p1.ice_cream_list {
		fmt.Printf("%v %v\n",k+1, v)

	}

}
