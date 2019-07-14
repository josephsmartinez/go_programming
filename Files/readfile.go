package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func exampleOne() {
	// Example One - wrapping the *File in *Scanner
	pFile, err := os.Open("test.txt")
	check(err)
	defer pFile.Close()

	reader := bufio.NewScanner(pFile)
	for reader.Scan() {
		token := reader.Text()
		fmt.Printf("Bytes: %v\n", len(token))
		fmt.Printf("String: %s\n", token)
	}
}

func exampleTwo() {
	// Example Two - brings the entire contents into memory (test.txt is 2MB)
	byteData, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Printf("%s\n", string(byteData))

	byteSize := (float64(len(string(byteData))) / math.Pow(2, 20))
	roundedUp := math.Ceil(byteSize)
	fmt.Printf("Byte Size: %.2f MB\n", roundedUp)

}

func exampleThree() {
	// Example Three - Reading from a buffer limit
	pFile, err := os.OpenFile("test.txt", os.O_RDONLY, 0644)
	check((err))
	reader := bufio.NewReader(pFile)
	defer pFile.Close()

	bufferSize := 1024
	buffer := make([]byte, bufferSize)
	// Read will fill buffer and return buff size and err (EOF)
	bytesRead, err := reader.Read(buffer)
	check(err)
	fmt.Printf("%s\n", string(buffer[:bytesRead]))

}

func exampleFour() {
	// Example Three - Reading from a buffer limit and looping
	pFile, err := os.Open("test.txt")
	check(err)
	defer pFile.Close()
	b1 := make([]byte, 1024)
	for {
		n1, err := pFile.Read(b1)
		if n1 > 0 {
			fmt.Printf("%s", string(b1[:n1]))
		}
		if err == io.EOF {
			break
		}
	}
}

func exampleFive() {
	// Example One - wrapping the *File with a buffered *Scanner
	pFile, err := os.Open("test.txt")
	check(err)
	defer pFile.Close()

	reader := bufio.NewScanner(pFile)
	buffer := make([]byte, 3344+1) // +1 for \0 explicit NUL terminator
	reader.Buffer(buffer, 0)

	for reader.Scan() {
		token := reader.Text()
		fmt.Printf("Bytes: %v\n", len(token))
		fmt.Printf("String: %s\n", token)
	}

}

func exampleSix() {
	// Example Six - Using all files in current path for reading
	files, err := ioutil.ReadDir(".")
	check(err)
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	fmt.Println(fileNames)

	for _, f := range fileNames {
		pfile, err := os.Open(f)
		check(err)
		b1 := make([]byte, 1024)
		for {
			n1, err := pfile.Read(b1)
			if n1 > 0 {
				fmt.Printf("%s", string(b1[:n1]))
			}
			if err == io.EOF {
				break
			}
		}
	}
}

func main() {

	//exampleOne()
	//exampleTwo()
	//exampleThree()
	//exampleFour()
	//exampleFive()
	//exampleSix()
}
