package main

import "fmt"

func main(){

	// Evaluating pointers
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	// Referenceing and dereferening pointers
	i, j := 42, 2701 
	p := &i 
	fmt.Println(*p) // dereferencing p 
	*p = 21
	fmt.Println(i) 

	p = &j
	*p = *p / 37    // dividing j through the pointer
	fmt.Println(j)

	// Function calls will pointers
	var a = f()
	fmt.Println(f() == f(), *a)

	// Function increments the value of the pointer
	v := 1
	incr(&v)
	fmt.Println(incr(&v)) // side effect increases 1 to 3

}

func f() *int{
	v := 1 
	return &v
}

func incr(p *int) int{
	*p++
	return *p
}