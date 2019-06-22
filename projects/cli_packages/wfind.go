package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Define flags
	pathPtr := flag.String("path", ".", "The path to search in (optional)")
	filePtr := flag.String("file", "", "The name of specific file (optional)")
	wordPtr := flag.String("word", "", "The word to look for")
	// Exit with code 1 if no args provided
	if len(os.Args) < 2 {
		fmt.Println("no args! Use -h or --help")
		os.Exit(1)
	}
	// Check if flags are clean to parse
	//add code here
	flag.Parse()
	fmt.Println("path", *pathPtr)
	fmt.Println("word", *wordPtr)
	fmt.Println("file", *filePtr)

	// Open files recursively to find file and word.
	file, err := os.Open("")
	check(err)
	defer file.Close()
	// Reading buffer size 3M
	buf := make([]byte, 32*1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Print(buf)
		}
		if err == io.EOF {
			break
		}
		check(err)
	}
}
