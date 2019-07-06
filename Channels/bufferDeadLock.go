package main

import (
	"fmt"
)

func main() {
	// To many channels for the buffer
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)

	c <- 43
	c <- 44
	fmt.Println(<-c)
}
