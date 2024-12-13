package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// SuccessRes represents the main structure with success information
type SuccessRes struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSON sends a JSON response to the client
func (h *Handler) WriteJSON(w http.ResponseWriter, status int) error {
	if h.Success == nil {
		return errors.New("no success response data")
	}

	out, err := json.Marshal(h.Success)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(out); err != nil {
		return err
	}

	h.Success = nil
	return nil
}
