package main

import "fmt"

const birthDay int = 1986

var currentYear int = 2019

func main() {
	yearsPast := birthDay

	for {
		fmt.Printf("%v,", yearsPast)
		yearsPast++
		if yearsPast > currentYear {
			break
		}
	}
}
