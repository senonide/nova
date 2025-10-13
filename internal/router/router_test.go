package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestNew(t *testing.T) {
	router := New()

	if router == nil {
		t.Fatal("Expected router to be created, got nil")
	}

	// Test that router has the expected type
	// Since router is already *gin.Engine, we just verify it's not nil
	// The type assertion is redundant here since New() returns *gin.Engine
}

func TestRouterMiddleware(t *testing.T) {
	router := New()

	// Add a test route to verify middleware is working
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestCORSMiddleware(t *testing.T) {
	router := New()

	// Add a test route
	router.GET("/cors-test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "cors test"})
	})

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "GET request with CORS headers",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "OPTIONS request",
			method:         http.MethodOptions,
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequestWithContext(context.Background(), tt.method, "/cors-test", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check CORS headers
			expectedOrigin := "*"
			if origin := w.Header().Get("Access-Control-Allow-Origin"); origin != expectedOrigin {
				t.Errorf("Expected Access-Control-Allow-Origin '%s', got '%s'", expectedOrigin, origin)
			}

			expectedMethods := "GET, POST, PUT, DELETE, OPTIONS"
			if methods := w.Header().Get("Access-Control-Allow-Methods"); methods != expectedMethods {
				t.Errorf("Expected Access-Control-Allow-Methods '%s', got '%s'", expectedMethods, methods)
			}

			expectedHeaders := "Content-Type, Authorization"
			if headers := w.Header().Get("Access-Control-Allow-Headers"); headers != expectedHeaders {
				t.Errorf("Expected Access-Control-Allow-Headers '%s', got '%s'", expectedHeaders, headers)
			}
		})
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	router := New()

	// Add a route that panics
	router.GET("/panic", func(_ *gin.Context) {
		panic("test panic")
	})

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/panic", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// The recovery middleware should handle the panic and return 500
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d after panic, got %d", http.StatusInternalServerError, w.Code)
	}
}
