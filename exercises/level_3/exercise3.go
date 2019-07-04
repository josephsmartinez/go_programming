package main

import (
	"fmt"
)

const birthDay int = 1986

var currentYear int = 2019

func main() {

	yearsPast := birthDay
	for yearsPast <= currentYear {
		fmt.Printf("%v,", yearsPast)
		yearsPast++
	}
}
