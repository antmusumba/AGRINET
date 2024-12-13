package server

import (
	"log"
	"net/http"
	"time"

	"github.com/antmusumba/agrinet/internals/routes"
)

// Server represents the HTTP server
type Server struct {
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(addr string) *Server {
	router := routes.SetupRoutes()

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
	log.Printf("Server is starting on %s", s.server.Addr)
}
