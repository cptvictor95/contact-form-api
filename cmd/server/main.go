package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Port for the server
	port := ":8000"

	fmt.Printf("Server is running on port %s\n", port)
	fmt.Println("Try visiting http://localhost:8000/health to check if the server is running")

	http.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write a JSON response
	response := `{"status": "healthy", "message": "Server is running"}`
	w.Write([]byte(response))
}