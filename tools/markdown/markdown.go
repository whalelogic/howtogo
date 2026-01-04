package main

import (
	"bytes"
	"os"
	"path/filepath"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)


// ConvertMarkdownToHTML converts a Markdown file to HTML format.
func ConvertMarkdownToHTML(inputPath, outputPath string) error {
	in, err := os.ReadDir(inputPath)
	if err != nil {
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	for _, file := range in {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}

		inputFilePath := filepath.Join(inputPath, file.Name())
		outputFileName := file.Name()[:len(file.Name())-3] + ".html"
		outputFilePath := filepath.Join(outputPath, outputFileName)
		content, err := os.ReadFile(inputFilePath)
		if err != nil {
			return err
	}
		var buf bytes.Buffer
		if err := md.Convert(content, &buf); err != nil {
			return err
		}
		if err := os.WriteFile(outputFilePath, buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil

		
}


func main() {
	inputPath := "/home/whaler/github/howtogo/content/markdown"
	outputPath := "/home/whaler/github/howtogo/content/html"
	
	if err := ConvertMarkdownToHTML(inputPath, outputPath); err != nil {
		panic(err)
	}
	fmt.Println("Markdown to HTML conversion completed successfully.")
}

