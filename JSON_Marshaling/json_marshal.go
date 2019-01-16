package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int8
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   34,
	}

	p2 := person{
		First: "Money",
		Last:  "Penny",
		Age:   29,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(bs))

}
