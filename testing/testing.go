package main

import (
	"fmt"
	"os"
)

// TestCreateFile used to create a file in /output directory
func TestCreateFile() {
	os.MkdirAll("/output", os.ModePerm)
	f, err := os.Create("output/testfile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("This is a test file.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	TestCreateFile()
}
