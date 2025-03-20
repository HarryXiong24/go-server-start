package middleware

import (
	"go-server-start/pkg/errors"
	"go-server-start/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key
// In production, this should be loaded from a secure configuration
var jwtSecret = []byte("your-secret-key")

// JWTClaims represents the claims in the JWT
type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a user
func GenerateToken(userID int64, name string) (string, error) {
	// Set expiration time - 24 hours from now
	expireTime := time.Now().Add(24 * time.Hour)

	// Create claims
	claims := JWTClaims{
		UserID: userID,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(userID),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString(jwtSecret)
}

// JWT returns a JWT auth middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(errors.NewUnauthorized("Authorization header is required", nil))
			c.Abort()
			return
		}

		// Check if the header format is valid
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.Error(errors.NewUnauthorized("Invalid authorization format, should be 'Bearer {token}'", nil))
			c.Abort()
			return
		}

		// Parse and validate token
		tokenString := parts[1]
		claims := &JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			logger.Sugar.Errorw("JWT token error", "error", err.Error())
			c.Error(errors.NewUnauthorized("Invalid or expired token", err))
			c.Abort()
			return
		}

		if !token.Valid {
			c.Error(errors.NewUnauthorized("Invalid token", nil))
			c.Abort()
			return
		}

		// Set user info to the context
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.Name)

		c.Next()
	}
}
