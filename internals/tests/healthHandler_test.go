package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/antmusumba/agrinet/internals/services"
)

func TestHandler_HealthHandler(t *testing.T) {
	handler := &handlers.Handler{
		AuthService:    &services.AuthService{},
		ProductService: &services.ProductService{},
	}

	req, err := http.NewRequest("GET", "/api/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler.HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", rec.Code, http.StatusOK)
	}

	resp := rec.Body.String()
	expectedResp := `{"status":"success","message":"Service is healthy and vibrating"}`
	if resp != expectedResp {
		t.Errorf("Handler returned wrong response: got %v want %v", resp, expectedResp)
	}
}
