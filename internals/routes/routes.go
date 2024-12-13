package routes

import (
	"net/http"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	return c.Handler(r)
}
