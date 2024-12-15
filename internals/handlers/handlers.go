package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/antmusumba/agrinet/internals/services"
	"github.com/antmusumba/agrinet/pkg"
)

// Handler represents the main handler structure that includes all services
type Handler struct {
	AuthService    *services.AuthService
	ProductService *services.ProductService
	Error          *ErrorRes
	Success        *SuccessRes
}

// NewHandler creates a new instance of Handler with combined services
func NewHandler(authService *services.AuthService, productService *services.ProductService) *Handler {
	return &Handler{
		AuthService:    authService,
		ProductService: productService,
	}
}

// HealthHandler handles health check requests
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := SuccessRes{
		Status:  "success",
		Message: "Service is healthy and vibrating",
	}

	h.Success = &response
	h.WriteJSON(w, http.StatusOK)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Invalid input",
		}
		h.Error = &errorRes
		h.WriteError(w, http.StatusBadRequest)
		return
	}

	if err := h.AuthService.Register(&user); err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Invalid input",
		}
		h.Error = &errorRes
		h.WriteError(w, http.StatusConflict)
		return
	}

	response := SuccessRes{
		Status:  "success",
		Message: "User registered successfully",
	}

	h.Success = &response
	h.WriteJSON(w, http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Invalid input",
		}
		h.Error = &errorRes

		h.WriteError(w, http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Login(credentials.Email, credentials.Password)
	if err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Invalid input",
		}
		h.Error = &errorRes

		h.WriteError(w, http.StatusUnauthorized)
		return
	}

	token, err := pkg.GenerateToken(user.ID)
	if err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Oops, something went wrong",
		}
		h.Error = &errorRes

		h.WriteError(w, http.StatusInternalServerError)
		return
	}

	response := SuccessRes{
		Status: "success",
		Data: map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.FirstName + " " + user.LastName,
			},
		},
	}

	h.Success = &response
	h.WriteJSON(w, http.StatusOK)
}

// CreateProduct handles product creation
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: "Invalid input",
		}
		h.Error = &errorRes
		h.WriteError(w, http.StatusBadRequest)
		return
	}

	err := h.ProductService.CreateProduct(&product)
	if err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: err.Error(),
		}
		h.Error = &errorRes
		h.WriteError(w, http.StatusInternalServerError)
		return
	}

	response := SuccessRes{
		Status:  "success",
		Message: "Product created successfully",
	}
	h.Success = &response
	h.WriteJSON(w, http.StatusCreated)
}

// ListProducts handles listing all products
func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.ListProducts()
	if err != nil {
		errorRes := ErrorRes{
			Status:  "error",
			Message: err.Error(),
		}
		h.Error = &errorRes
		h.WriteError(w, http.StatusInternalServerError)
		return
	}

	response := SuccessRes{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    products,
	}

	h.Success = &response
	h.WriteJSON(w, http.StatusOK)
}


