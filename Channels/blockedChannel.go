package main

import (
	"fmt"
)

func main() {
	ca := make(chan int)
	ca <- 42
	fmt.Println(<-ca)
}
