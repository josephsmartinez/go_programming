package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
)

func main() {
	this_year := 2019
	fmt.Printf("%d %d %d %d\n", a, b, c, d)
	fmt.Printf("%d %d %d %d\n", this_year+a, this_year+b, this_year+c, this_year+d)

}
