package routes

import (
	"net/http"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/antmusumba/agrinet/internals/repositories"
	"github.com/antmusumba/agrinet/internals/services"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Router represents the main router structure
type Router struct {
	MuxRouter *mux.Router
	Handler   *handlers.Handler
}

// NewRouter initializes a new router with dependencies
func NewRouter(userRepo repositories.UserRepo, productRepo repositories.ProductRepo) *Router {
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	return &Router{
		MuxRouter: mux.NewRouter(),
		Handler:   handlers.NewHandler(authService, productService),
	}
}

// SetupRoutes configures all the routes for the application
func (r *Router) SetupRoutes() http.Handler {
	r.MuxRouter.HandleFunc("/api/health", r.Handler.HealthHandler).Methods("GET")
	r.MuxRouter.HandleFunc("/api/auth/register", r.Handler.Register).Methods("POST")
	r.MuxRouter.HandleFunc("/api/auth/login", r.Handler.Login).Methods("POST")

	r.MuxRouter.HandleFunc("/api/products", r.Handler.CreateProduct).Methods("POST")
	r.MuxRouter.HandleFunc("/api/products", r.Handler.ListProducts).Methods("GET")

	// TODO: Add product routes
	//
	// r.muxRouter.HandleFunc("/api/products", r.handler.GetProducts).Methods("GET")
	// r.muxRouter.HandleFunc("/api/products/{id}", r.handler.UpdateProduct).Methods("PUT")
	// r.muxRouter.HandleFunc("/api/products/{id}", r.handler.DeleteProduct).Methods("DELETE")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	return c.Handler(r.MuxRouter)
}
