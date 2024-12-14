package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"project/handlers"
	"project/pkg"
	"testing"

	"project/mocks" // Mocked dependencies

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mocks.NewMockAuthService(ctrl)
	handler := &handlers.Handler{
		AuthService: mockAuthService,
	}

	// Mock the token generator function
	mockTokenGenerator := func(userID string) (string, error) {
		if userID == "valid-user-id" {
			return "valid-token", nil
		}
		return "", errors.New("token generation failed")
	}
	pkg.GenerateToken = mockTokenGenerator

	// Test cases
	testCases := []struct {
		name           string
		requestBody    map[string]string
		mockSetup      func()
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid login",
			requestBody: map[string]string{
				"email":    "testuser@example.com",
				"password": "validpassword",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().
					Login("testuser@example.com", "validpassword").
					Return(&pkg.User{
						ID:        "valid-user-id",
						Email:     "testuser@example.com",
						FirstName: "Test",
						LastName:  "User",
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "success",
				"data": map[string]interface{}{
					"token": "valid-token",
					"user": map[string]interface{}{
						"id":    "valid-user-id",
						"email": "testuser@example.com",
						"name":  "Test User",
					},
				},
			},
		},
		{
			name: "Invalid credentials",
			requestBody: map[string]string{
				"email":    "testuser@example.com",
				"password": "wrongpassword",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().
					Login("testuser@example.com", "wrongpassword").
					Return(nil, errors.New("invalid credentials"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]interface{}{
				"status":  "error",
				"message": "Invalid input",
			},
		},
		{
			name: "Token generation failure",
			requestBody: map[string]string{
				"email":    "testuser@example.com",
				"password": "validpassword",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().
					Login("testuser@example.com", "validpassword").
					Return(&pkg.User{
						ID:        "invalid-user-id",
						Email:     "testuser@example.com",
						FirstName: "Test",
						LastName:  "User",
					}, nil)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status":  "error",
				"message": "Oops, something went wrong",
			},
		},
		{
			name: "Invalid request payload",
			requestBody: map[string]string{
				"email": "missing-password",
			},
			mockSetup:      func() {}, // No mock setup needed
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status":  "error",
				"message": "Invalid input",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock setup for this test case
			tc.mockSetup()

			// Prepare the request
			requestBody, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Simulate the response
			recorder := httptest.NewRecorder()
			handler.Login(recorder, req)

			// Validate the status code
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			// Validate the response body
			var responseBody map[string]interface{}
			json.NewDecoder(recorder.Body).Decode(&responseBody)
			assert.Equal(t, tc.expectedBody, responseBody)
		})
	}
}
