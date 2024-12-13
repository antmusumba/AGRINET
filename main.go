package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antmusumba/agrinet/internals/database"
	"github.com/antmusumba/agrinet/internals/repositories"
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
	db, err := database.InitDB("agrinet.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories with database connection
	userRepo := repositories.NewUserRepo(db)

	// Create and start server
	srv := server.NewServer(GetPort(), userRepo)
	srv.Start()
}
