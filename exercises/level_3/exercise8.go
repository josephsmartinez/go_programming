package main

import "fmt"

var (
	banana string = "banana"
	orange string = "orange"
	apple  string = "apple"
	potato string = "potato"
	mongo  string = "mongo"
	pear   string = "pear"
)

func main() {

	fruits := []string{banana, orange, apple, potato, mongo, pear}
	for _, fruit := range fruits {
		switch {
		case fruit == "banana":
			fmt.Println("banana")
		case fruit == "orange":
			fmt.Println("orange")
		case fruit == "apple":
			fmt.Println("apple")
		case fruit == "potato":
			fmt.Println("potato")
		case fruit == "mongo":
			fmt.Println("mongo")
		case fruit == "pear":
			fmt.Println("pear")
		}
	}
}
