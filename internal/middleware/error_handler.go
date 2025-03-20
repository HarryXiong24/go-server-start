package middleware

import (
	"go-server-start/pkg/errors"
	"go-server-start/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler returns a middleware for handling errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Check if there are any errors
		if len(c.Errors) == 0 {
			return
		}

		// Get the last error
		err := c.Errors.Last().Err

		// Check if it's a custom error
		var statusCode int
		var errorResponse gin.H

		// Check for custom error types
		if appErr, ok := err.(*errors.AppError); ok {
			statusCode = appErr.StatusCode
			errorResponse = gin.H{
				"error":  appErr.Message,
				"code":   appErr.Code,
				"status": statusCode,
			}
		} else {
			// Default error handling
			statusCode = http.StatusInternalServerError
			errorResponse = gin.H{
				"error":  "Internal Server Error",
				"status": statusCode,
			}
		}

		// Log the error
		logger.Sugar.Errorw("Request error",
			"error", err.Error(),
			"path", c.Request.URL.Path,
			"status", statusCode,
		)

		// Return error response
		c.JSON(statusCode, errorResponse)
	}
}
