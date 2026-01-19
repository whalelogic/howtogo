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
// 		log.Fatalf("Failed to set trusted proxies: %v", err)
// 	}
// 
// 	// Define routes
// 	r.GET("/health", h.HealthCheck)
// 	r.POST("/analytics/event", h.RecordEvent)
// 	r.GET("/analytics/stats", h.GetStats)
// 
// 	// Start the server
// 	log.Printf("Starting server on port %s", port)
// 	if err := r.Run(":" + port); err != nil {
// 		log.Fatalf("Failed to run server: %v", err)
// 	}
// }



func main() {
	TestCreateFile()
}
