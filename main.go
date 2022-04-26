package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Program started")

	var filenames []string

	// read list of files from directory
	dir := "./original"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory")
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jhf") {
			filenames = append(filenames, file.Name())
		}
	}

	// print list of files
	fmt.Println("Files to be processed:")
	for _, filename := range filenames {
		fmt.Println(filename)
	}

}
