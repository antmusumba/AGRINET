package handlers

import (
	"net/http"
)

// Handler represents the main handler structure
type Handler struct {
	Success *SuccessRes
}

// NewHandler creates a new instance of Handler
func NewHandler() *Handler {
	return &Handler{}
}

// HealthHandler handles health check requests
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.Success = &SuccessRes{
		Status:  "success",
		Message: "Service is healthy",
	}
	h.WriteJSON(w, http.StatusOK)
}
