package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleShortenURL(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	r := gin.Default()
	r.POST("/shorten", handleShortenURL)

	// Test cases
	tests := []struct {
		name       string
		payload    ShortenRequest
		wantStatus int
	}{
		{
			name: "Valid URL",
			payload: ShortenRequest{
				URL: "https://www.example.com",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Empty URL",
			payload: ShortenRequest{
				URL: "",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Perform request
			r.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantStatus == http.StatusOK {
				// Parse response
				var response ShortenResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.NotEmpty(t, response.ShortURL)
			}
		})
	}
}

func TestGenerateShortCode(t *testing.T) {
	// Test multiple generations to ensure uniqueness
	codes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		code, err := generateShortCode()
		assert.NoError(t, err)
		assert.Len(t, code, 8)
		assert.False(t, codes[code], "Generated duplicate code: %s", code)
		codes[code] = true
	}
} 