package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antmusumba/agrinet/internals/handlers"
	"github.com/antmusumba/agrinet/internals/models"
	"github.com/stretchr/testify/assert"
)

// Stub implementation of ProductService
type StubProductService struct {
	ListProductsFunc func() ([]*models.Product, error)
}

func (s *StubProductService) ListProducts() ([]*models.Product, error) {
	return s.ListProductsFunc()
}

func TestListProductsHandler(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		mockSetup      func() *StubProductService
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid products list",
			mockSetup: func() *StubProductService {
				return &StubProductService{
					ListProductsFunc: func() ([]*models.Product, error) {
						return []*models.Product{
							{ID: "1", Name: "Product 1", Price: 10.0},
							{ID: "2", Name: "Product 2", Price: 20.0},
						}, nil
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status":  "success",
				"message": "Products retrieved successfully",
				"data": []interface{}{
					map[string]interface{}{"id": "1", "name": "Product 1", "price": 10.0},
					map[string]interface{}{"id": "2", "name": "Product 2", "price": 20.0},
				},
			},
		},
		{
			name: "Error fetching products",
			mockSetup: func() *StubProductService {
				return &StubProductService{
					ListProductsFunc: func() ([]*models.Product, error) {
						return nil, errors.New("failed to fetch products")
					},
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status":  "error",
				"message": "failed to fetch products",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup the handler with the mock ProductService
			handler := &handlers.Handler{
				ProductService: tc.mockSetup(),
			}

			// Simulate the HTTP request
			req, _ := http.NewRequest(http.MethodGet, "/products", nil)
			recorder := httptest.NewRecorder()
			handler.ListProducts(recorder, req)

			// Assert the status code
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			// Assert the response body
			var responseBody map[string]interface{}
			json.NewDecoder(recorder.Body).Decode(&responseBody)
			assert.Equal(t, tc.expectedBody, responseBody)
		})
	}
}
