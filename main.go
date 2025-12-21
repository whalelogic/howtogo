package main

import (
	"fmt"
	"io"
	"net/http"
)

type server struct {
	addr string
	port int
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "200 OK. You are healthy.\n")
	if err != nil {
		fmt.Println("Error writing response:", err)
	} else {
		fmt.Println("Health check responded with OK")
	}
}

func main() {

	srv := &server{
		addr: "localhost",
		port: 8080,
	}
	http.HandleFunc("/health", CheckHealth)

	address := fmt.Sprintf("%s:%d", srv.addr, srv.port)
	fmt.Printf("Starting server at %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
