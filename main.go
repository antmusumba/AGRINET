package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// healthChecker returns a simple health check response
func healthChecker(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "success",
		Message: "The server is breezing in a healthy way!",
		Data: map[string]string{
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/health", healthChecker).Methods("GET")

	// Create CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, c.Handler(r)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
