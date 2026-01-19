package main

import (
	"fmt"
	"os"
	"path/filepath"
	"context"

	"github.com/whalelogic/howtogo/views/components"


)


func main() {

	f, err := os.Create(filepath.Join("output", "button.html"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	templComponent := components.Button("Click Me", "http://example.com").Render(context.Background(), f)
	fmt.Println("Static page generated successfully.")
	fmt.Println(templComponent)


}
