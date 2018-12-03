/*Echo4 
Package flag implement command line parsing
https://godoc.org/flag
Package strings manipulate UTF-8 encoding
https://godoc.org/strings
*/
package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}