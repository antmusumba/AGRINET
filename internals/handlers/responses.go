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

// ErrorRes represents the main structure with error information
type ErrorRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// WriteError prepares and sends an error response in JSON format
func (h *Handler) WriteError(w http.ResponseWriter, status int) error {
	if h.Error == nil {
		h.Error = &ErrorRes{
			Status:  "error",
			Message: "Internal server error",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(h.Error); err != nil {
		return err
	}

	return nil
}

// ReadJSON decodes a JSON request body into a data structure
func (h *Handler) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("expected application/json content-type")
	}

	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1048576))
	if err := dec.Decode(data); err != nil {
		return errors.New("invalid JSON body")
	}

	if err := dec.Decode(&struct{}{}); err == nil {
		return errors.New("unexpected extra data in JSON body")
	}
	return nil
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
	h.Error = nil
	return nil
}
