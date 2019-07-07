package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (s SecretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p Person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func randomFunction(h human) {
	switch h.(type) {
	case Person:
		fmt.Println("I was passed into randomFunction", h.(Person).first)
	case SecretAgent:
		fmt.Println("I was passed into randomFunction", h.(SecretAgent).first)
	}
	fmt.Println("I was passed into randomFunction", h)
}

type hotdog int

func main() {
	sa1 := SecretAgent{
		Person: Person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := SecretAgent{
		Person: Person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := Person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	randomFunction(sa1)
	randomFunction(sa2)
	randomFunction(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
