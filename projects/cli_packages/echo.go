package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Newline n and Separator sep
var n = flag.Bool("n", false, "omit training lines")
var sep = flag.String("s", " ", "Separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(os.Args, *sep))
	if !*n {
		fmt.Println()
	}
}
