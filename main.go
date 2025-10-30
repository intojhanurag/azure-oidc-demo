package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from GitHub Actions + Azure OIDC (Go version)! 🚀")
}

func main() {
	// Use the PORT from environment (Azure sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local testing
	}

	http.HandleFunc("/", handler)
	log.Printf("✅ Server running on port %s\n", port)
	// Prepend "0.0.0.0:" to the port to explicitly listen on all network interfaces
	address := "0.0.0.0:" + port
	log.Fatal(http.ListenAndServe(address, nil))
}

