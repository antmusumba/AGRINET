package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/antmusumba/agrinet/internals/services"
	"github.com/antmusumba/agrinet/pkg"
)

// Handler represents the main handler structure
type Handler struct {
	service *services.AuthService
	Error   *ErrorRes
	Success *SuccessRes
}

// NewHandler creates a new instance of Handler
func NewHandler(service *services.AuthService) *Handler {
	return &Handler{service: service}
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

	if err := h.service.Register(&user); err != nil {
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

	user, err := h.service.Login(credentials.Email, credentials.Password)
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
