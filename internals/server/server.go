package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/antmusumba/agrinet/internals/repositories"
	"github.com/antmusumba/agrinet/internals/routes"
)

// Server represents the HTTP server
type Server struct {
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(addr string, userRepo repositories.UserRepo, productRepo repositories.ProductRepo) *Server {
	router := routes.NewRouter(userRepo, productRepo).SetupRoutes()

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

// handleShutdown handles the shutdown of the server
func (s *Server) handleShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown did not complete in %v : %v", 15*time.Second, err)
		if err := s.server.Close(); err != nil {
			log.Fatalf("Could not stop server gracefully : %v", err)
		}
	}
	log.Println("Server shutdown completed")
}

// Start starts the server and handles graceful shutdown
func (s *Server) Start() {
	serverErrors := make(chan error, 1)
	commandChan := make(chan string, 1)

	// Start the server
	go func() {
		log.Printf("Server is starting on %s", s.server.Addr)
		serverErrors <- s.server.ListenAndServe()
	}()

	// Start command reader
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("server> ")
			command, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading command: %v", err)
				continue
			}
			command = strings.TrimSpace(strings.ToLower(command))
			commandChan <- command
		}
	}()

	// Handle OS signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown signal
	for {
		select {
		case err := <-serverErrors:
			log.Fatalf("Error starting server: %v", err)

		case sig := <-shutdown:
			log.Printf("Received shutdown signal: %v", sig)
			s.handleShutdown()
			return

		case cmd := <-commandChan:
			switch cmd {
			case "exit":
				log.Println("Received exit command")
				s.handleShutdown()
				return
			case "help":
				fmt.Println("Available commands:")
				fmt.Println("  exit - Shutdown the server")
				fmt.Println("  help - Show help message")
			default:
				fmt.Printf("Unknown command: %s\nType 'help' for available commands\n", cmd)
			}
		}
	}
}
