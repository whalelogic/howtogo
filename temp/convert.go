package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	file, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	outputDir := "html"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	for _, f := range file {
		fmt.Println(f.Name())
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			inputName := f.Name()
			outputName := strings.TrimSuffix(inputName, ".md") + ".html"

			src := filepath.Clean(inputName)
			dst := filepath.Join(outputDir, outputName)

			cmd := exec.Command("pandoc", src, "-o", dst)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error converting %s to %s: %v\n", inputName, outputName, err)
			} else {
				fmt.Printf("Converted %s to %s.\n", inputName, dst)

			}
		}
	}

	fmt.Println("Done.")
}
