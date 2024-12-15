package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antmusumba/agrinet/internals/services"
	"github.com/antmusumba/agrinet/internals/handlers"
)

func TestHandler_CreateProduct(t *testing.T) {
	handler := &handlers.Handler{
		AuthService:    &services.AuthService{},
		ProductService: &services.ProductService{},
	}

	req, err := http.NewRequest("POST", "/api/products", strings.NewReader(`{"name":"Test Product","price":10,"description":"This is a test product"}`))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler.CreateProduct(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", rec.Code, http.StatusCreated)
	}

	resp := rec.Body.String()
	expectedResp := `{"status":"success","message":"Product created successfully"}`
	if resp != expectedResp {
		t.Errorf("Handler returned wrong response: got %v want %v", resp, expectedResp)
	}
}
