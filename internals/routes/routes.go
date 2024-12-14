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
	muxRouter *mux.Router
	handler   *handlers.Handler
}

// NewRouter initializes a new router with dependencies
func NewRouter(userRepo repositories.UserRepo, productRepo repositories.ProductRepo) *Router {
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	return &Router{
		muxRouter: mux.NewRouter(),
		handler:   handlers.NewHandler(authService, productService),
	}
}

// SetupRoutes configures all the routes for the application
func (r *Router) SetupRoutes() http.Handler {
	r.muxRouter.HandleFunc("/api/health", r.handler.HealthHandler).Methods("GET")
	r.muxRouter.HandleFunc("/api/paymentgateway", r.handler.PaymentHandler).Methods("GET")
	r.muxRouter.HandleFunc("/api/auth/register", r.handler.Register).Methods("POST")
	r.muxRouter.HandleFunc("/api/auth/login", r.handler.Login).Methods("POST")

	r.muxRouter.HandleFunc("/api/products", r.handler.CreateProduct).Methods("POST")
	r.muxRouter.HandleFunc("/api/products", r.handler.ListProducts).Methods("GET")

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

	return c.Handler(r.muxRouter)
}
