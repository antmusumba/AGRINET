package server

import (
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Response represents a standard API response
type Response struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

// SetupRoutes configures all the routes for the application
func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Return the router with CORS middleware
	return c.Handler(r)
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "success",
		Message: "Gracefully shutdown the server",
		Data: map[string]string{
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Server represents the HTTP server
type Server struct {
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(addr string) *Server {
	router := SetupRoutes()

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		server: srv,
	}
}

// Start starts the server and handles graceful shutdown
func (s *Server) Start() {
	// Start the server
	log.Printf("Server is starting on %s", s.server.Addr)
}
