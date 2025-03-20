package main

import (
	"fmt"
	"go-server-start/internal/config"
	"go-server-start/internal/routers"
	"go-server-start/pkg/database"
	"go-server-start/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		fmt.Println("Failed to load configuration:", err)
		os.Exit(1)
	}

	// Initialize logger
	if err := logger.Init(config.AppConfig.Logger.Level); err != nil {
		fmt.Println("Failed to initialize logger:", err)
		os.Exit(1)
	}

	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.Mode)
	server := gin.Default()

	// Initialize database connection
	if err := database.Init(); err != nil {
		logger.Sugar.Errorf("Failed to initialize database: %v", err)
		os.Exit(1)
	}

	// Initialize routes
	router := server.Group("/")
	routers.Init(router)

	// Start server
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	logger.Sugar.Infof("Server started on port %s", port)
	if err := server.Run(port); err != nil {
		logger.Sugar.Errorf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
