// FUNCTIONS USED
// Open()
// file.Read()

package main

import (
	"fmt"
	"log"
	"os"
)

func readWithBuffer() {

	// Returns a file object
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)

	}

	// Read from file object with buffer size
	var totalReadBytes = 0
	for {
		dataBuffer := make([]byte, 100)
		count, err := file.Read(dataBuffer)

		if count == 0 {
			fmt.Printf("bytes read: %d \n", totalReadBytes)
			os.Exit(0)
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q", dataBuffer[:count])
		totalReadBytes += count
	}
}

func main() {

	readWithBuffer()

}
