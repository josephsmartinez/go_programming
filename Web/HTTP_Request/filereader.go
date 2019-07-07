package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	fileName string
)

const (
	bufferSize int = (5 * 1024)
)

var w = bufio.NewWriterSize(new(os.File), bufferSize)

w


func print() {
	fmt.Print("")
}
