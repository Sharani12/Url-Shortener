package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func handleShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate short code
	shortCode, err := generateShortCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short code"})
		return
	}

	// Store in Redis
	ctx := context.Background()
	err = redisClient.Set(ctx, shortCode, req.URL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	// Increment counter
	urlCounter.Inc()

	// Return response
	c.JSON(http.StatusOK, ShortenResponse{
		ShortURL: c.Request.Host + "/" + shortCode,
	})
}

func handleRedirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code is required"})
		return
	}

	// Retrieve from Redis
	ctx := context.Background()
	originalURL, err := redisClient.Get(ctx, shortCode).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	// Increment counter
	urlCounter.Inc()

	// Redirect to original URL
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:8], nil
} 