package main

import (
	"fmt"
	"os"
  "log"
)

func main() {
	file, err := os.Open("text.txt") // For read access.
	if err != nil {
		log.Fatal(err)

	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
