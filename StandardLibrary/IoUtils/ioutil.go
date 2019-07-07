package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// File file data
type File struct {
	FileName string
	FileSize string
	FileMode int
	FileMod  string
	IsDir    string
	Perm     string
}

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var fileSlice []File
	for _, f := range files {

		fileJSON := File{
			FileName: f.Name(),
			FileSize: fmt.Sprintf("%.3f", float64(float64(f.Size())/1024)),
			FileMode: int(f.Mode()),
			FileMod:  f.ModTime().String(),
			IsDir:    fmt.Sprintf("%t", f.IsDir()),
			Perm:     string(f.Mode().Perm()),
		}
		fileSlice = append(fileSlice, fileJSON)

		// fmt.Printf("File Name: %v\nFile Size: %.3f MB\nFile Mode: %d\nLast Mod: %v\nIs Directory: %v\nUnix Permission Bits: %v\n\n", f.Name(), float64(float64(f.Size())/1024), f.Mode(), f.ModTime().String(), f.IsDir(), f.Mode().Perm())
	}
	fmt.Println(fileSlice)

	filesJSON, err := json.Marshal(fileSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(filesJSON))
}
