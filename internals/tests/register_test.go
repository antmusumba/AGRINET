package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/antmusumba/agrinet/internals/services"
)

func TestHandler_Register(t *testing.T) {
	handler := &handlers.Handler{
		AuthService:    &services.AuthService{},
		ProductService: &services.ProductService{},
	}

	req, err := http.NewRequest("POST", "/api/auth/register", strings.NewReader(`{"email":"test@example.com","password":"password"}`))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler.Register(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", rec.Code, http.StatusCreated)
	}

	resp := rec.Body.String()
	expectedResp := `{"status":"success","message":"User registered successfully"}`
	if resp != expectedResp {
		t.Errorf("Handler returned wrong response: got %v want %v", resp, expectedResp)
	}
}
