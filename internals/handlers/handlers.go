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
	authService    *services.AuthService
	productService *services.ProductService
	Error          *ErrorRes
	Success        *SuccessRes
}

// NewHandler creates a new instance of Handler with combined services
func NewHandler(authService *services.AuthService, productService *services.ProductService) *Handler {
	return &Handler{
		authService:    authService,
		productService: productService,
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

func (h *Handler) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentDetails struct {
		Amount      int    `json:"amount"`
		PhoneNumber string `json:"phone"`
	}

	// Decode incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&paymentDetails); err != nil {
		h.WriteError(w, http.StatusBadRequest)
		return
	}

	// Call the STK Push service to process the payment
	resp, err := services.ProcessStkPush(paymentDetails.PhoneNumber, paymentDetails.Amount)
	if err != nil {
		h.WriteError(w, http.StatusBadRequest)
		return
	}

	// Send the success response
	h.Success = &SuccessRes{
		Status:  "success",
		Message: "Payment processed successfully",
		Data:    resp,
	}

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

	if err := h.authService.Register(&user); err != nil {
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

	user, err := h.authService.Login(credentials.Email, credentials.Password)
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

	err := h.productService.CreateProduct(&product)
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
	products, err := h.productService.ListProducts()
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
