package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Error getting hostname", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hostname: %s\n", hostname)
}

func main() {
	// Define a command-line flag for the port
	port := flag.Int("port", 8080, "Port for the server to listen on")
	flag.Parse()

	// Register the handler for the root endpoint
	http.HandleFunc("/", handler)

	// Start the server on the specified port
	fmt.Printf("Server is running on http://localhost:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
