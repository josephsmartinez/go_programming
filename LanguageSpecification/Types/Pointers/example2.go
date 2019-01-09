package main

import (
	"fmt"
)

type person struct{
	name string
	age int
	address string

}

func changeMe(p *person){
	p.address = "123 sw 12 street"
	(*p).name = "Joe Mart"
}

func main() {

	p1 := person{
		name : "James Bond",
		age : 34,
		address: "007 Street",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
