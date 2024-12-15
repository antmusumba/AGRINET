package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/antmusumba/agrinet/internals/services"

)

func TestHandler_Login(t *testing.T) {
	handler := &handlers.Handler{
		AuthService:    &services.AuthService{},
		ProductService: &services.ProductService{},
	}

	req, err := http.NewRequest("POST", "/api/auth/login", strings.NewReader(`{"email":"test@example.com","password":"password"}`))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler.Login(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", rec.Code, http.StatusOK)
	}

	resp := rec.Body.String()
	expectedResp := `{"status":"success","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE1MTYyMzkwMjJ9.fjDTqRncapSB94hkJvQZ70ITZK0KuYf7PEnlCBwoUO8","user":{"id":1,"email":"test@example.com","name":"Test User"}}`
	if resp != expectedResp {
		t.Errorf("Handler returned wrong response: got %v want %v", resp, expectedResp)
	}
}
