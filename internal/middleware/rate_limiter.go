package middleware

import (
	"go-server-start/pkg/logger"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// SimpleRateLimiter is a basic rate limiter implementation
type SimpleRateLimiter struct {
	sync.Mutex
	requests map[string][]time.Time
	window   time.Duration
	max      int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(window time.Duration, max int) *SimpleRateLimiter {
	return &SimpleRateLimiter{
		requests: make(map[string][]time.Time),
		window:   window,
		max:      max,
	}
}

// RateLimiter returns a middleware for limiting request rates
// It limits based on client IP address
func RateLimiter(window time.Duration, max int) gin.HandlerFunc {
	limiter := NewRateLimiter(window, max)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		// Check if the request should be limited
		if !limiter.Allow(ip) {
			logger.Sugar.Warnw("Rate limit exceeded", "ip", ip)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Allow checks if a request from the given IP is allowed
func (r *SimpleRateLimiter) Allow(ip string) bool {
	r.Lock()
	defer r.Unlock()

	now := time.Now()

	// Clean up old requests
	if _, exists := r.requests[ip]; exists {
		var validRequests []time.Time
		for _, t := range r.requests[ip] {
			if now.Sub(t) <= r.window {
				validRequests = append(validRequests, t)
			}
		}
		r.requests[ip] = validRequests
	}

	// Check if the rate limit is exceeded
	if len(r.requests[ip]) >= r.max {
		return false
	}

	// Add the current request
	r.requests[ip] = append(r.requests[ip], now)
	return true
}
