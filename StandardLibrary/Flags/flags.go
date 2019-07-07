package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// definedflags
	var pModNumber = flag.Int("chmod", 0644, "enter a four bit value for the file permissions. ie. 0644, 0755, etc")

	// no args
	if len(os.Args) < 2 {
		fmt.Println("no args! Use -h or --help")
		os.Exit(1)
	}

	// parse flage args
	flag.Parse()

	// file name from args
	filenames := os.Args[2:]
	flagArg := os.Args[1:2]

	if flagArg[0] != "-h" && flagArg[0] != "--help" && !strings.Contains(flagArg[0], "--chmod=") {
		fmt.Println("no flag args! Use -h or --help")
		os.Exit(1)

	}

	// check if the args length are valid
	argLen := len(filenames)
	if argLen > 1 || argLen < 1 {
		fmt.Println("incorrent number of file args!")
		// take a list of file(s)
		fmt.Println("You entered the following args:")
		for v := range filenames {
			fmt.Printf("{ type: %T, file: %v\n", filenames[v], filenames[v])
		}
		os.Exit(1)
	}

	// chnage file permission
	if err := os.Chmod(filenames[0], os.FileMode(*pModNumber)); err != nil {
		log.Fatal(err)
	}

}
