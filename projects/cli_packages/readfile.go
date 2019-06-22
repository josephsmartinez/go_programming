package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func read(files []string) {

	for _, f := range files {
		f, err := os.Open(f)
		check(err)
		b1 := make([]byte, 1024)
		for {
			n1, err := f.Read(b1)
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

	files, err := ioutil.ReadDir(".")
	check(err)
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	fmt.Println(fileNames)

	// SMALL UNIT TEST
	//files := []string{"test.txt", "test1.txt"}
	//read(fileNames)
}
