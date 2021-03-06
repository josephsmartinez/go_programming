package main

import "fmt"

// Person : Capital names are avalible outside the package level
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Lower case structs are for this package only
type employee struct {
	// The type does not need a field name
	Person
	jobTitle    string
	companyName string
	salary      int
}

// NOTE: namespace with field "size"
type tires struct {
	size int
	make string
}

type car struct {
	tires tires
	make  string
	model string
	year  int
	miles int
}

func main() {

	// Empty initialization
	person := new(Person)
	fmt.Println(person)

	// Creating a struture using field names
	person1 := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	// Creating a structure without field names
	person2 := Person{"Gold", "Finger", 42}
	fmt.Println(person2)

	// Is-A Relationship
	employee1 := employee{
		// Initializing does require a field name
		Person: Person{
			firstName: "Jeff",
			lastName:  "Bezos",
			age:       55,
		},
		jobTitle:    "CEO",
		companyName: "Amazon",
		salary:      10000000000,
	}

	fmt.Printf("First Name:%v\nLast Name: %v\nAge: %v\n",
		person1.firstName, person1.lastName, person1.age)

	fmt.Printf("First Name:%v\nLast Name: %v\nAge: %v\nJob Title: %v\nCompany: %v\nSalary: %v\n", employee1.firstName, employee1.lastName, employee1.age, employee1.jobTitle, employee1.companyName, employee1.salary)

	// Has-A Relationship
	car1 := car{
		tires: tires{
			size: 10,
			make: "goodyear",
		},
		make:  "honda",
		model: "civic",
		year:  2018,
		miles: 18000,
	}

	fmt.Printf("Make: %v\nModel: %v\nYear: %v\nMiles: %v\nTiress: %v %v\n", car1.make, car1.model, car1.year, car1.miles, car1.tires.make, car1.tires.size)

	// Anonymous Structures
	iphone := struct {
		model string
		price int
	}{
		model: "5s",
		price: 100,
	}
	fmt.Printf("%T\n", iphone)
}
