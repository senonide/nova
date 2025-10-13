package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/senonide/nova/internal/router"
)

func TestMain(m *testing.M) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	m.Run()
}

func setupTestRouter() *gin.Engine {
	r := router.New()
	RegisterRoutes(r)
	return r
}

func TestGetNova(t *testing.T) {
	router := setupTestRouter()
	response := makeRequest(t, router, http.MethodGet, "/api/v1/nova")

	expectedMessage := "Go Nova!"
	if message, ok := response["message"].(string); !ok || message != expectedMessage {
		t.Errorf("Expected message '%s', got '%v'", expectedMessage, response["message"])
	}

	expectedStatus := "success"
	if status, ok := response["status"].(string); !ok || status != expectedStatus {
		t.Errorf("Expected status '%s', got '%v'", expectedStatus, response["status"])
	}
}

func TestGetHealth(t *testing.T) {
	router := setupTestRouter()
	response := makeRequest(t, router, http.MethodGet, "/api/v1/health")

	expectedStatus := "healthy"
	if status, ok := response["status"].(string); !ok || status != expectedStatus {
		t.Errorf("Expected status '%s', got '%v'", expectedStatus, response["status"])
	}

	expectedTime := "up"
	if time, ok := response["time"].(string); !ok || time != expectedTime {
		t.Errorf("Expected time '%s', got '%v'", expectedTime, response["time"])
	}
}

func makeRequest(t *testing.T, router *gin.Engine, method, path string) map[string]any {
	t.Helper()

	req, err := http.NewRequestWithContext(context.Background(), method, path, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	return response
}

func TestGetRoot(t *testing.T) {
	router := setupTestRouter()
	response := makeRequest(t, router, http.MethodGet, "/")

	validateRootResponse(t, response)
}

func validateRootResponse(t *testing.T, response map[string]any) {
	t.Helper()

	expectedMessage := "Welcome to Nova API"
	if message, ok := response["message"].(string); !ok || message != expectedMessage {
		t.Errorf("Expected message '%s', got '%v'", expectedMessage, response["message"])
	}

	expectedVersion := "1.0.0"
	if version, ok := response["version"].(string); !ok || version != expectedVersion {
		t.Errorf("Expected version '%s', got '%v'", expectedVersion, response["version"])
	}

	validateEndpoints(t, response)
}

func validateEndpoints(t *testing.T, response map[string]any) {
	t.Helper()

	endpoints, ok := response["endpoints"].(map[string]any)
	if !ok {
		t.Error("Expected endpoints object in response")
		return
	}

	expectedNova := "/api/v1/nova"
	if nova, ok := endpoints["nova"].(string); !ok || nova != expectedNova {
		t.Errorf("Expected nova endpoint '%s', got '%v'", expectedNova, endpoints["nova"])
	}

	expectedHealth := "/api/v1/health"
	if health, ok := endpoints["health"].(string); !ok || health != expectedHealth {
		t.Errorf("Expected health endpoint '%s', got '%v'", expectedHealth, endpoints["health"])
	}
}

func TestNotFoundEndpoint(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/nonexistent", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/api/v1/nova", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestCORSHeaders(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/v1/nova", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

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
}

func TestOptionsRequest(t *testing.T) {
	router := setupTestRouter()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodOptions, "/api/v1/nova", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d for OPTIONS request, got %d", http.StatusNoContent, w.Code)
	}
}
