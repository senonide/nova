package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes(r *gin.Engine) {
	// API v1 routes
	v1 := r.Group("/api/v1")
	v1.GET("/nova", GetNova)
	v1.GET("/health", GetHealth)

	// Root endpoint
	r.GET("/", GetRoot)
}

// GetNova handles the main nova endpoint.
func GetNova(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Go Nova!",
		"status":  "success",
	})
}

// GetHealth handles health check requests.
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"time":   "up",
	})
}

// GetRoot handles root path requests.
func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Nova API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"nova":   "/api/v1/nova",
			"health": "/api/v1/health",
		},
	})
}
