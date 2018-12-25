package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // ErrHelp is the error returned if the -help or -h flag is invoked but no such flag is defined.
    flag.Usage = func() {
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("    example7 file1 file2 ...\n")
        flag.PrintDefaults()
    }
    flag.Parse()
}
