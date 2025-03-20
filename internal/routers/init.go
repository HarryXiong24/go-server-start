package routers

import (
	"go-server-start/internal/handlers"
	"go-server-start/internal/middleware"
	"go-server-start/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Init initializes all routes
func Init(router *gin.RouterGroup) {
	logger.Sugar.Info("Initializing API routes")

	// Add global middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	// // Public API group with rate limiting
	// api := router.Group("/api")
	// api.Use(middleware.RateLimiter(time.Minute, 60)) // 60 requests per minute
	// {
	// 	// User routes (public for now)
	// 	users := api.Group("/users")
	// 	{
	// 		users.GET("/:id", handlers.GetUserInfo)
	// 		users.GET("", handlers.GetUserInfo)
	// 	}
	// }

	// // Protected API group (requires authentication)
	// // Currently not applied to any routes
	// protected := router.Group("/api/auth")
	// {
	// 	// JWT Authentication middleware
	// 	protected.Use(middleware.JWT())

	// 	// Protected routes will be added here when needed
	// 	// No routes for now as there's no need for authentication yet
	// }

	// Legacy route that needs to be migrated
	router.GET("/get-user-info", handlers.GetUserInfo)
}
