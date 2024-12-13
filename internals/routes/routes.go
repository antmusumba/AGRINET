package routes

import (
	"net/http"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Router represents the main router structure
type Router struct {
	muxRouter *mux.Router
	handler   *handlers.Handler
}

// NewRouter initializes a new router with dependencies
func NewRouter() *Router {
	return &Router{
		muxRouter: mux.NewRouter(),
		handler:   handlers.NewHandler(),
	}
}

// SetupRoutes configures all the routes for the application
func (r *Router) SetupRoutes() http.Handler {
	r.muxRouter.HandleFunc("/api/health", r.handler.HealthHandler).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	return c.Handler(r.muxRouter)
}
