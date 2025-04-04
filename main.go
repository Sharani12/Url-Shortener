package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	redisClient *redis.Client
	urlCounter  prometheus.Counter
)

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	// Initialize Prometheus metrics
	urlCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "url_shortener_requests_total",
		Help: "Total number of URL shortening requests",
	})
	prometheus.MustRegister(urlCounter)
}

func main() {
	r := gin.Default()

	// Add middleware
	r.Use(rateLimiter())
	r.Use(gin.Recovery())

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	r.POST("/shorten", handleShortenURL)
	r.GET("/:shortCode", handleRedirect)

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simple rate limiting implementation
		// In production, you might want to use a more sophisticated approach
		time.Sleep(100 * time.Millisecond) // Simulate rate limiting
		c.Next()
	}
} 