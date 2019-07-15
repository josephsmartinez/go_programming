package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exampleOne() {

	reader := bufio.NewReader(os.Stdin)
	for {
		stdin, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading console")
			continue
		} else if strings.ToLower(strings.Trim(stdin, "\n")) == "q" {
			os.Exit(0)
		}
		// Echo Results
		fmt.Printf("Read: %v", stdin)
	}

}

func main() {
	exampleOne()

}
