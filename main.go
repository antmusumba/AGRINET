package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antmusumba/agrinet/internals/server"
)

// GetPort returns the port from environment variable or default value
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("No PORT environment variable found, using default: %s", port)
	}
	return fmt.Sprintf(":%s", port)
}

// main is the entry point of the application
func main() {
	srv := server.NewServer(GetPort())

	// Start the server
	srv.Start()
}
