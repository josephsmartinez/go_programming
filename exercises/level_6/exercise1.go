package main

import "fmt"

func foo() int {
	return 32
}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}

func main() {
	foo := foo()
	bar, bar1 := bar()
	fmt.Printf("%v %v %v\n", foo, bar, bar1)
}
